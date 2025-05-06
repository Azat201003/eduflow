<script setup>
    import { ref } from 'vue'
    const user = ref({})
    var authed = ref(false)
    const token = localStorage.getItem("token")
    if (token == '') {

    } else {
        fetch("http://localhost:8080/my-user/", {
            method: "GET",
            cache: "no-cache",
            mode: "cors",
            credentials: "omit",
            headers: {
                "Authorization": "Bearer " + token,
            },
        })
            .then((response) => {
                if (response.status != 200) {
                    // document.querySelector("#error-message").innerHTML = "some error";
                    console.log(response.status)
                    return
                } else {
                    response.json().then((json) => {
                        user.value = json
                        console.log(user.value)
                        console.log(1)
                        authed.value = true
                        // document.querySelector("#error-message").innerHTML = "";
                    })
                }
            })
    }
    function quit() {
        localStorage.removeItem("token")
        location.reload()
    }
</script>

<template>
    <div class="container">
        <div class="profile-view" v-if="authed">
            <div class="row">
                <div class="col-4"></div>
                <div class="col-4">
                    <h1>{{ user.username }}</h1>
                    <p>is_staff = <i>{{ user.is_staff }}</i></p>
                </div>
                <div class="col-4"></div>
            </div>
            <div class="row">
                <div class="col-5"></div>
                <button class="quit col-2 btn-primary btn" @click="quit">quit</button>
                <div class="col-5"></div>
            </div>
        </div>
        <div class="no-profile-view" v-else>
            You are don't authed
        </div>
    </div>
</template>

<style>
.profile-view {
    min-height: 200px;
    text-align: center;
    vertical-align: middle;
    margin-top: 200px;
}
.no-profile-view {
    min-height: 200px;
    text-align: center;
    vertical-align: middle;
    margin-top: 200px;
}
</style>
