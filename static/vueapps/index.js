const indexApp = Vue.createApp (
    {
        data () {
            return {

            }
        },
        methods: {
            handleEvent() {
                console.log("Reached")
            }
        }
    }
)
indexApp.mount("#indexApp")