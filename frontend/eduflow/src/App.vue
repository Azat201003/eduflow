<script setup>
import 'bootstrap'
import 'bootstrap/dist/css/bootstrap.min.css';
import Navbar from '@/components/Navbar.vue';
</script>

<template>
  <head>
    <title>abeme</title>
  </head>
  <main>
    <Navbar/>
    <!-- <nav>
      <ul>
        <li><RouterLink to="/">Home</RouterLink></li>
        <li><RouterLink to="/auth/sign-in">Sign in</RouterLink></li>
        <li><RouterLink to="/auth/sign-up">Sign up</RouterLink></li>
      </ul>
    </nav> -->
    <RouterView /> 
  </main>
</template>

<style scoped>
/* @import "~bootstrap/scss/bootstrap"; */
header {
  line-height: 1.5;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }
}
</style>
