<template>
  <div class="container">
    <h1>Register</h1>
    <div class="row">
      <div class="col-md-4">
        <form @submit.prevent="onSubmit">
          <div class="form-group">
            <label for="username">Username</label>
            <input type="text" v-model="registerForm.username" class="form-control" v-bind:class="{'is-invalid': registerForm.usernameError}" id="username" placeholder="">
            <div v-show="registerForm.usernameError" class="invalid-feedback">{{ registerForm.usernameError }}</div>
          </div>
          <div class="form-group">
            <label for="email">Email address</label>
            <input type="email" v-model="registerForm.email" class="form-control" v-bind:class="{'is-invalid': registerForm.emailError}" id="email" aria-describedby="emailHelp" placeholder="">
            <small v-if="!registerForm.emailError" id="emailHelp" class="form-text text-muted">We'll never share your email with anyone else.</small>
            <div v-show="registerForm.emailError" class="invalid-feedback">{{ registerForm.emailError }}</div>
          </div>
          <div class="form-group">
            <label for="password">Password</label>
            <input type="password" v-model="registerForm.password" class="form-control" v-bind:class="{'is-invalid': registerForm.passwordError}" id="password" placeholder="">
            <div v-show="registerForm.passwordError" class="invalid-feedback">{{ registerForm.passwordError }}</div>
          </div>
          <button type="submit" class="btn btn-primary">Register</button>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import store from '../store'

export default {
  name: 'Register',
  data: function () {
    return {
        registerForm: {
            username: '',
            email: '',
            password: '',
            submitted: false,
            errors: 0,
            usernameError: null,
            emailError: null,
            passwordError: null,
        }
    }
  },
    methods: {
        onSubmit (e) {
           this.registerForm.submitted = true
           this.registerForm.errors = 0

           if (!this.registerForm.username) {
               this.registerForm.errors ++
               this.registerForm.usernameError = 'Username required.'
           } else {
               this.registerForm.usernameError = null
           }

           if (!this.registerForm.email) {
               this.registerForm.errors ++
               this.registerForm.emailError = 'Email required.'
           } else {
               this.registerForm.emailError = null
           }

           if (!this.registerForm.password) {
               this.registerForm.errors ++
               this.registerForm.passwordError = 'Password required.'
           } else {
               this.registerForm.passwordError = null
           }

           if (this.registerForm.errors > 0) {
               return false
           }

           const path = 'http://127.0.0.1:8888/register'
           let param = new URLSearchParams()
           param.append("username", this.registerForm.username)
           param.append("email", this.registerForm.email)
           param.append("password", this.registerForm.password)
           

           axios.post(path, param)
           .then((res)=>{
             if (res.data.code == 200) {
               store.setNewAction()
               this.$router.push('/login')
             } else {
               if (res.data.code == 20007) {
                 this.registerForm.emailError = 'Email already exists.'
               } else if (res.data.code == 20006) {
                 this.registerForm.usernameError = 'Username already exists.'
               } else {
                 this.registerForm.passwordError = 'Invalid register.'
               }
             }
           })
        }
    }
}
</script>

<style scoped>

</style>
