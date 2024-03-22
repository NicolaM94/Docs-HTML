const registervue = Vue.createApp ({
    data () {
        return {
            buttonstatus :false,
            errormessage:"",
            name:"",
            surname:"",
            email:"",
            emailagain:"",
            password:"",
            passwordagain:""

        }
    },
    methods: {
        checkName () {
            if (this.name.length == 1) {
                this.errormessage = "Il nome deve contenere più di un carattere"
                return false
            } else {
                this.errormessage = ""
                return true
            }
        },
        checkSurname () {
            if (this.surname.length == 1) {
                this.errormessage = "Il cognome deve contenere più di un carattere"
                return false
            } else {
                this.errormessage = ""
                return true
            }
        },
        checkEmail() {
            if (!this.email.includes("@")) {
                this.errormessage = "Indirizzo email non valido"
                return false
            } else {
                this.errormessage = ""
                return true
            }
        },
        checkAgainEmail () {
            let equity = this.email === this.emailagain
            if (!equity) {
                this.errormessage = "Gli indirizzi email non coincidono"
                return false
            } else {
                this.errormessage = ""
                return true
            }
        },
        checkPasswordAgain () {
            let equity = this.password === this.passwordagain
            if (!equity) {
                this.errormessage = "Le password non coincidono"
                this.buttonstatus = false
                return false
            } else {
                this.errormessage = ""
                this.buttonstatus = true
                return true
            }
        }        
    }
})

registervue.mount("#registervue")