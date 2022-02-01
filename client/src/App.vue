<template>
  <v-app>
    <v-app-bar
      app
      dark
    >
      <div class="d-flex align-center">
        <router-link to="/" style="text-decoration: none; color: white;">
          <v-row>
            <v-img class="mr-2 ml-16" src="./assets/logo-2.png" width="42px"></v-img>
            <h2>TRing</h2>
          </v-row>
        </router-link>
      </div>
    </v-app-bar>

    <v-main>
      <router-view/>
    </v-main>

    <footer class="" style="background-color: #1e1e1e;">
      <div align="center" class="my-3">
        <p class="footer-text" style="color: white;">&copy; {{ new Date().getFullYear() }} YutoSekiguchi & ShinyaMoriyasu</p>
      </div>
    </footer>
  </v-app>
</template>

<script>

export default {
  name: 'App',

  data: () => ({
    //
  }),
};
</script>