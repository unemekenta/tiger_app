import axios from 'axios'

export const loginMixin = {
  created () {
    this.$_loginMixin_loginCheck();
  },
  data () {
    return {
      signedIn: false,
      jwtData: "",
    }
  },
  methods: {
    async $_loginMixin_loginCheck() {
      var data = "";
      if (window.$cookies.isKey('uuid')) {
        await axios
          .get(process.env.VUE_APP_API_BASE_URL + "/api/login_check", {
            params: {
              key: window.$cookies.get('uuid')
            },
          })
          .then((res) => {
            this.jwtData = res.data;
          })
          .catch(() => {
            // チェックなので未ログインでも問題はないのでエラーハンドリングなし
            return;
          });

        await axios
          .post(
            process.env.VUE_APP_API_BASE_URL + '/auth/api/authCheck',
            data,
            {
              headers: {
                'Authorization': 'Bearer ' + this.jwtData,
              },
            }
          )
          .then(() => {
            this.signedIn = true;
            return;
          })
          .catch(() => {
            // チェックなので未ログインでも問題はないのでエラーハンドリングなし
            this.signedIn = false;
            return;
          });
      } else {
        // チェックなので未ログインでも問題はないのでエラーハンドリングなし
        return;
      }
    },

    $_loginMixin_signOut() {
      window.$cookies.remove("uuid");
      this.signedIn = false;
      this.$router.push("/signin");
    },
  }
}