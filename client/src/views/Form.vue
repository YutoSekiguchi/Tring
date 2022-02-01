<template>
  <v-container>

    <v-dialog
        v-model="dialog"
        max-width="500px"
      >
      <v-card>
        <v-card-title>
          名前を記入してください
        </v-card-title>
        <v-card-text>
          <v-text-field
            v-model="userName"
            label="ex. 山田花子"
            autofocus
            @keyup.enter="moveNowPage()"
          ></v-text-field>
        </v-card-text>
        <v-card-actions>
          <v-btn
            color="primary"
            text
            @click="moveToShow()"
          >
            閉じる
          </v-btn>
          <v-spacer></v-spacer>
          <v-btn
            color="primary"
            dark
            @click="moveNowPage()"
          >
            登録
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <h1 class="font-weight-bold" style="font-size: 12rem; color: #92CBDD;" align="center"><span style="color: #F8BA60;">T</span>Ring</h1>

    <div style="background-color: rgb(245, 245, 255);" class="mt-12">
      <v-row>
        <v-col cols="12" align="center" v-if="this.user"><h4 class="font-weight-bold mt-4">ようこそ<span style="color: #5C6BC0; font-size: 24px;">{{ this.user.Name }}</span>さん</h4></v-col>
        <v-col cols="12" align="center"><h3 class="font-weight-bold">休憩時間（分）を入力してください</h3></v-col>
        <v-col cols="4"></v-col>
        <v-col cols="4">
          <v-text-field
            v-model="minute"
            type="number"
            label="休憩時間を入力"
            class="num-form"
            autofocus
            @keyup.enter="onSubmit()"
          ></v-text-field>
        </v-col>
        <v-col cols="4"></v-col>
        <v-col cols="12" align="center">
          <v-btn
            color="blue-grey darken-2"
            @click="onSubmit()"
          >
            <p class="mt-2" style="color: white;">Enter</p>
          </v-btn>
        </v-col>
      </v-row>
    </div>
  </v-container>
</template>

<script>

  export default {
    name: 'Form',
    data() {
      return {
        dialog: false,
        userName: null,
        userList: null,
        user: null,
        minute: 10,
        time: null,
        isTime: null,
      }
    },
    methods: {
      async onSubmit() { // 確定ボタン押した時の処理
        if ((Number(this.minute < 0))) {
          alert("自然数を入力してください");
          this.minute = 10;
        }
        else {
          var created_date = new Date(); // 現在の時刻
          var finish_date = new Date();

          await finish_date.setMinutes(finish_date.getMinutes() + Number(this.minute)) //終了予定時刻

          var created_at = this.dateToString(created_date);
          var finish_at  = this.dateToString(finish_date); 

          await this.axios({
            method: 'POST',
            url: `/times`,
            data: {
              uid: Number(this.user.ID),
              time: Math.floor(Number(this.minute)),
              finishAt: String(finish_at),
              createdAt: String(created_at),
            }
          }).then((res) => {
            console.log(res.data);
          }).catch((e) => {
            console.log(e);
          })
          
          this.$router.push(`/show`);
        }
      },
      async moveNowPage() { // 名前入力後ページ
        await this.axios({
          method: 'POST',
          url: `/users`,
          data: {
            name: this.userName,
            card: this.$route.params.card,
          }
        }).then((res) => {
          console.log(res.data);
          this.dialog = false;
        }).catch((e) => {
          console.log(e);
        });
        await this.getUser();
        this.user = this.userList[0];

        await this.axios({
          method: 'GET',
          url: `/times/time/${this.user.ID}`
        }).then((res) => {
          this.time = res.data;
        }).catch((e) => {
          console.log(e);
        });
      },
      moveToShow() { // 閉じるボタンを押した時
        this.dialog = false;
        this.$router.push(`/show`);
      },
      async getUser() {
        await this.axios({
          method: 'GET',
          url: `/users/card/${this.$route.params.card}`,
        }).then((res) => {
          console.log(res.data);
          this.userList = res.data
        }).catch((e) => {
          console.log(e);
        });
      },
      dateToString(date) {
        const strYear = String(date.getFullYear()).padStart(4, '0')
        const strMonth = String(date.getMonth()+1).padStart(2, '0')
        const strDate = String(date.getDate()).padStart(2, '0')
        const strHour = String(date.getHours()).padStart(2, '0')
        const strMin = String(date.getMinutes()).padStart(2, '0')
        const strSec = String(date.getSeconds()).padStart(2, '0')
        return strYear + '-' +  strMonth + '-' + strDate + ' ' + strHour + ':' + strMin + ':' + strSec
      }
    },
    async created() {
      await this.getUser();

      if (!this.userList) {
        this.dialog = true;
      } else {
        this.user = this.userList[0];

        await this.axios({
          method: 'GET',
          url: `/times/time/${this.user.ID}`
        }).then((res) => {
          this.time = res.data;
        }).catch((e) => {
          console.log(e);
        });

        if (this.time) {
          await this.axios({
            method: 'DELETE',
            url: `/times/usertime/delete/${this.user.ID}`
          }).catch((e) => {
            console.log(e);
          })
          await this.axios({
            method: 'GET',
            url: `/times/list`
          }).then((res) => {
            this.isTime = res.data;
          }).catch((e) => {
            console.log(e);
          });
          if (this.isTime) {
            this.$router.push(`/show`);
          } else {
            this.$router.push(`/`);
          }
        }
      }
    }
  }
</script>