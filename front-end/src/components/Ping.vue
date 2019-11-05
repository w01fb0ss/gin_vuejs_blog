<template>
    <div class="container">
        <button type="button" class="btn btn-primary">{{ msg }}</button>
    </div>
</template>

<script>
import axios from "axios"

export default {
    name: "Ping",
    data() {
        return {
            msg: ''
        }
    },
    methods: {
        getMessage () {
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