<template>
    <div class="container">
        <alert 
            v-if="sharedState.is_new"
            v-bind:variant="alertVariant"
            v-bind:message="alertMessage">
        </alert>
        <h1>Sign In</h1>
        <div class="row">
            <div class="col-md-4">
                <form @submit.prevent="onSubmit">
                <div class="form-group">
                    <label for="username">Username</label>
                    <input type="text" v-model="loginForm.username" class="form-control" v-bind:class="{'is-invalid': loginForm.usernameError}" id="username" placeholder="">
                    <div v-show="loginForm.usernameError" class="invalid-feedback">{{ loginForm.usernameError }}</div>
                </div>
                <div class="form-group">
                    <label for="password">Password</label>
                    <input type="password" v-model="loginForm.password" class="form-control" v-bind:class="{'is-invalid': loginForm.passwordError}" id="password" placeholder="">
                    <div v-show="loginForm.passwordError" class="invalid-feedback">{{ loginForm.passwordError }}</div>
                </div>
                <button type="submit" class="btn btn-primary">Sign In</button>
                </form>
            </div>
        </div>
        <br>
        <p>New User? <router-link to="/register">Click to Register!</router-link></p>
        <p>
            Forgot Your Password?
            <a href="#">Click to Reset It</a>
        </p>
    </div>
</template>

<script>
import axios from 'axios'
import Alert from './Alert'
import store from '../store'



export default {
    name: "Login",
    components: {
        alert: Alert,
    },
    data () {
        return {
           sharedState: store.state,
           alertVariant: 'info',
           alertMessage: 'Congratulations, you are now a registered user !',
           loginForm: {
               username: '',
               password: '',
               submitted: false,
               errors: 0,
               usernameError: null,
               passwordError: null
           }
        }
    },
    methods: {
        onSubmit (e) {
            this.loginForm.submitted = true
            this.loginForm.errors = 0

            if (!this.loginForm.username) {
                this.loginForm.errors ++
                this.loginForm.usernameError = 'Username required.'
            } else {
                this.loginForm.errors = null
            }

            if (!this.loginForm.password) {
                this.loginForm.errors ++
                this.loginForm.usernameError = 'Password required.'
            } else {
                this.loginForm.errors = null
            }

            if (this.loginForm.errors > 0 ) {
                return false
            }

            const path = 'http://127.0.0.1:8888/auth'
            let param = new URLSearchParams()
            param.append("username", this.loginForm.username)
            param.append("password", this.loginForm.password)
            axios.post(path, param)
            .then((res) => {
                if (res.data.code == 200) {
                    window.localStorage.setItem('token', res.data.data.token)
                    store.resetNotNewAction()
                    store.loginAction()

                    if (typeof this.$route.query.redirect == 'undefined') {
                        this.$router.push('/ping')
                    } else {
                        this.$router.push(this.$router.query.redirect)
                    }
                } else {
                    this.loginForm.usernameError = 'Invalid username or password.'
                    this.loginForm.passwordError = 'Invalid username or password.'
                }
                
            })
        }
    }
}
</script>