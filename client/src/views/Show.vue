<template>
  <div>
    <v-row class="mt-3">
      <v-col v-for="n in timesList.length" :key="n" cols="12" sm="6" md="4" lg="3">
        <Timer 
          :timeId="timesList[n-1].ID"
          :name="timesList[n-1].Name" 
          :time="timesList[n-1].Time" 
          :color="colorsList[n-1 % 8]"
          :finishAt="timesList[n-1].FinishAt" 
          :createdAt="timesList[n-1].CreatedAt"
          :haveTime="calcTime(timesList[n-1].FinishAt)"
        />
      </v-col>
    </v-row>
  </div>
</template>

<script>
  import Timer from "../components/Timer.vue";

  export default {
    name: 'Show',
    data() {
      return {
        num: 2,
        timesList: [],
        colorsList: [
          "red accent-1",
          "cyan accent-2",
          "green accent-3",
          "amber accent-4",
          "purple lighen-1",
          "lime accent-2",
          "pink lighten-1",
          "teal darken-1",
        ],
      }
    },
    components: {
      Timer,
    },
    methods: {
      calcTime(finishTime) {
      var ft = Date.parse(finishTime);
      var nt = new Date();
      
      var diffTime = ft - nt;
      var diffSecond = Math.floor(diffTime / (1000));
      return diffSecond
    }
    },
    async created() {
      await this.axios({
        method: 'GET',
        url: `/times/list`
      }).then((res) => {
        this.timesList = res.data;
      }).catch((e) => {
        console.log(e);
      });
      if (!this.timesList) {
        this.$router.push("/");
      }
    },
  }
</script>