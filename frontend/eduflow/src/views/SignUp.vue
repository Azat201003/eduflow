<script lang="ts" setup>
    function sign_up(event: SubmitEvent) {
        event.preventDefault()
        const formData = new FormData(<HTMLFormElement>event.target);
        fetch("http://localhost:8080/auth/sign-up/", {
            method: "POST",
            body: formData
        }).then((response) => {
            response.json().then((json) => {
                if (!json.ok){
                    document.querySelector("#error-message").innerHTML = json.message;
                    return
                }
                console.log(response.status)
                document.querySelector("#error-message").innerHTML = "";
                response.json().then((json) => {
                    localStorage.removeItem("token")
                    localStorage.setItem("token", json["token"])
                })
                window.location.href = "/"
            })
        })
    }
    window.onload = function (e) {
        document.getElementById("register").addEventListener('submit', sign_up)
    }
</script>
<template>
    <div class="container">
        <div class="row">
            <div class="col-4"></div>
            <div class="col-4 d-grid">
                <form id="register">
                    <div class="mb-3">
                        <label class="form-label">Username</label>
                        <input type="username" class="form-control" id="username-input" aria-describedby="emailHelp" name="username">
                    </div>
                    <div class="mb-3">
                        <label class="form-label">Password</label>
                        <input type="password" class="col form-control" id="password-input" name="password">
                    </div>
                    <div id="error-message"></div>
                    <button type="submit" class="btn btn-primary btn-block btn-lg mb-4">Sign up</button>
                </form>
            </div>
        </div>
    </div>
</template>
