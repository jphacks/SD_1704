package sessions

import (
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/jphacks/SD_1704/model"
)

// store is session store based on gorilla/sessions.
// This is singleton for wiki app.
var store = sessions.NewCookieStore([]byte("secretkey"))

const SESSION_KEY = "session"
const SESSION_ID = "sessionId"

func Get(r *http.Request, key string) (*sessions.Session, error) {
	return store.Get(r, key)
}

func Save(r *http.Request, w http.ResponseWriter, session *sessions.Session) error {
	return store.Save(r, w, session)
}

// Clear removes the given session
func Clear(r *http.Request, w http.ResponseWriter, session *sessions.Session) error {
	session.Options.MaxAge = -1
	return Save(r, w, session)
}

func AllClear(r *http.Request, w http.ResponseWriter) error {
	sessionWrite, err := Get(r, SESSION_KEY)
	if err != nil {
		return err
	}
	Clear(r, w, sessionWrite)
	return nil
}

func ManageSession(db *sql.DB, userId int64, r *http.Request, w http.ResponseWriter) error {
	// UUIDを生成する
	uuid := uuid.New().String()

	//セッションをDBに保存
	//セッションがすでにある場合は更新

	res, err := model.SessionExistsByUserId(db, userId)
	if err != nil {
		return err
	}

	if res {
		//セッションがある場合
		err = model.UpdateSessionByUserId(db, uuid, userId)
		if err != nil {
			return err
		}
	} else {
		//セッションがない場合
		err = model.CreateSession(db, uuid, userId)
		if err != nil {
			return err
		}
	}

	//クッキーにセッションIDをつける。
	sessionWrite, err := Get(r, SESSION_KEY)
	if err != nil {
		return err
	}
	sessionWrite.Values[SESSION_ID] = uuid
	Save(r, w, sessionWrite)
	return err
}

func GetUserIdBySession(db *sql.DB, r *http.Request) (*int64, error) {
	// セッションを貰う
	sessionWrite, err := Get(r, SESSION_KEY)
	if err != nil {
		return nil, err
	}
	sessionUuid := sessionWrite.Values[SESSION_ID]
	// セッション(uuid)が存在しない場合
	if sessionUuid == nil {
		return nil, nil
	}
	// データベースセッションからユーザーIDを調べる
	userId, err := model.GetUserIdByUuid(db, sessionUuid.(string))
	if err != sql.ErrNoRows && err != nil {
		return nil, err
	}
	return userId, nil
}

func IsShouldRedirect(db *sql.DB, w http.ResponseWriter, r *http.Request, userId *int64) (user model.User, err error, shouldRedirect bool) {
	if userId == nil {
		return model.User{}, nil, true
	}

	user, err = model.GetUserByUserId(db, *userId)
	if err != nil {
		return model.User{}, err, false
	}

	return user, nil, false
}
