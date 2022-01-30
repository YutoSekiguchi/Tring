<template>
  <div>
    <v-dialog
        v-model="dialog"
        max-width="500px"
      >
      <v-card>
        <v-card-title>
          {{ name }}さん
        </v-card-title>
        <v-card-text>
          休憩時間を削除しますか？
        </v-card-text>
        <v-card-actions>
          <v-btn
            color="primary"
            text
            @click="dialog=false"
          >
            閉じる
          </v-btn>
          <v-spacer></v-spacer>
          <v-btn
            color="error"
            dark
            @click="deleteTime(timeId)"
          >
            削除
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <v-card 
      align="center"
      class="ml-1 mr-1"
      @click="dialog=true"
      v-if="this.t > 0"
    >
      <v-card-title alien="center" class="pb-0">
        <p class="font-weight-bold" style="border-bottom: 2px solid rgb(110, 110, 180);"><span style="font-size: 1.5rem;">{{ name }}</span>さん</p>
      </v-card-title>
      <v-progress-circular
        :rotate="-90"
        :size="320"
        :width="35"
        :value="((this.time*60 - this.t)*100/(this.time*60))"
        :color="color"
        class="mb-5"
      >
      <h1 class="font-weight-bold" align="center" style="font-size: 5rem;">
        {{ parseInt(this.time - (this.time*60 - this.t) / 60) + 1 }}分
      </h1>
      </v-progress-circular>
    </v-card>

    <v-card 
      align="center"
      class="ml-1 mr-1"
      @click="dialog=true"
      v-if="this.t <= 0"
      dark
    >
      <v-card-title alien="center" class="pb-0">
        <p class="font-weight-bold" style="border-bottom: 2px solid rgb(110, 110, 180);"><span style="font-size: 1.5rem;">{{ name }}</span>さん</p>
      </v-card-title>
      <v-progress-circular
        :rotate="-90"
        :size="320"
        :width="25"
        :value="((this.time*60 - this.t)*100/(this.time*60))"
        :color="color"
        class="mb-5"
      >
      <h1 class="font-weight-bold" align="center" style="font-size: 5rem;">
        {{ parseInt(this.time - (this.time*60 - this.t) / 60) }}分
      </h1>
      </v-progress-circular>
    </v-card>
  </div>
</template>

<script>
export default {
  name: 'Timer',
  data() {
    return {
      dialog: false,
      finishTime: null,
      nowTime: null,
      diffTime: null,
      t: 0,
    }
  },
  props: {
    timeId: {
      type: Number
    },
    name: {
      type: String
    },
    time: {
      type: Number
    },
    color: {
      type: String
    },
    finishAt: {
      type: String
    },
    createdAt: {
      type: String
    },
    haveTime: {
      type: Number
    }
  },
  computed: {
  },
  methods: {
    async deleteTime(id) {
      await this.axios({
        method: 'DELETE',
        url: `/times/delete/${id}`
      }).catch((e) => {
        console.log(e);
      });
      this.dialog = false;
      window.location.reload();
    }
  },
  created() {
    this.t = this.haveTime;
    setInterval(() => {
      if (this.t > -600) {
        this.t--;
      } else {
        this.deleteTime(this.timeId)
      }
    }, 1000)
  }
}
</script>