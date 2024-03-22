const loginvuehandler = Vue.createApp({
    data () {
        return {
            emailfield : "",
            emailerror : true
        }
    },
    methods: {
        checkemailerror () {
            if (!this.emailfield.includes("@")) {
                this.emailerror = true
            } else {
                this.emailerror = false
            }
        }
    }
})

loginvuehandler.mount("#loginvue")