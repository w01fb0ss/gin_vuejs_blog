<template>
    <div class="container">
        <alert 
            v-if="sharedState.is_authenticated"
            v-bind:variant="alertVariant"
            v-bind:message="alertMessage">
        </alert>
        <h1>Ping</h1>
        <button type="button" class="btn btn-primary">{{ msg }}</button>
    </div>
</template>

<script>
import axios from "axios"
import store from '../store'
import Alert from './Alert'

export default {
    name: "Ping",
    components: {
        alert: Alert,
    },
    data() {
        return {
            msg: '',
            sharedState: store.state,
            alertVariant: 'info',
            alertMessage: 'Congratulations, you has login success !',
        }
    },
    methods: {
        getMessage () {
            console.log(this.sharedState)
            const path = 'http://localhost:8888/api/ping'
            let token = window.localStorage.getItem("token")
            axios.get(path, {
                params: {
                    token: token,
                }
            })
            .then((res) => {
                this.msg = res.data;
            })
            .catch((error) => {
                console.log(error)
            })
        }
    },
    created() {
        this.getMessage()
    }
}
</script>