<template>
  <div>
      <div>username: <input type="text"  v-model="username"  /></div>
      <div>password: <input type="text"  v-model="password"  /></div>
      <div><input type="button" value="LOGIN" v-on:click="onLogin()" ></div>
      {{msg}}
  </div>
</template>


<script lang="ts">
import Login from './scripts/Login';
export default Login;
</script>
