Vue.component("common-menu", {
  props: ["current"],
  template: `
    <div class="ui secondary pointing menu">
      <a class="item" href="./index.html" v-bind:class="{active: current=='index'}">
      トップ
      </a>
      <a class="item" href="./userinfo-update.html" v-bind:class="{active: current=='userinfo'}">
        プロフィール
      </a>
      <a class="item" href="./point.html" v-bind:class="{active: current=='point'}">
        ポイント
      </a>
      <a class="item" href="./board.html" v-bind:class="{active: current=='board'}">
        掲示板
      </a>
      <div class="right menu">
        <a class="ui item" v-on:click="logout">
          ログアウト
        </a>
      </div>
    </div>`,
methods: {
  logout: function () {
    localStorage.removeItem('token');
    location.href = "./login.html";
  }
}
});
