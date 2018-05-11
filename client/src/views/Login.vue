<template>
  
  <div class="ui middle aligned center aligned grid" >
    <div class="top"></div>
    <div class="row" style="margin:8px;" >
      <div class="six wide tablet six wide computer sixteen wide column">
        <!-- head -->
        
        <!-- form -->
        <div class="ui large form" >
          <div class="ui stacked segment">
            <h2 class="ui teal image header">
              <img src="https://semantic-ui.com/examples/assets/images/logo.png" class="image">
              <div class="content">Log-in to your account</div>
            </h2>
            <div class="field">
              <div class="ui left icon input">
                <i class="user icon"></i>
                <input type="text" name="usernamme" placeholder="Username" v-model="username" >
              </div>
            </div>
            <div class="field">
              <div class="ui left icon input">
                <i class="lock icon"></i>
                <input type="password" name="password" placeholder="Password" v-model="password" >
              </div>
            </div>
              <div><sui-button   class="ui fluid large teal submit button" primary v-on:click="onLogin()" >Login</sui-button></div>
              <div>{{msg}}</div>
            </div>  
          </div>
      </div>
    </div>
  </div>
</template>


<script lang="ts">
import Login from "./scripts/Login";
export default Login;
</script>

<style>
body{
  height: 100%;
  background-color:#2c3e50;
}
.top{
  min-height: 150px;
}
</style>
