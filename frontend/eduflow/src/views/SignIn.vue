<script lang="ts" setup>
    function sign_in(event: SubmitEvent) {
        event.preventDefault()
        const formData = new FormData(<HTMLFormElement>event.target);

        fetch("http://localhost:8080/auth/sign-in/", {
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
                localStorage.removeItem("token")
                localStorage.setItem("token", json.data.token)
                window.location.href = "/"
            })
        })
    }
    window.onload = function (e) {
        document.getElementById("login").addEventListener('submit', sign_in)
    }
</script>

<template>
    <div class="container">
        <div class="row">
            <div class="col-4"></div>
            <div class="col-4 d-grid">
                <form id="login" target="_blank">
                    <div class="mb-3">
                        <label class="form-label">Username</label>
                        <input type="username" class="form-control" id="username-input" name="username" aria-describedby="emailHelp" required>
                    </div>
                    <div class="mb-3">
                        <label class="form-label">Password</label>
                        <input type="password" class="col form-control" id="password-input" name="password" required>
                    </div>
                    <div id="error-message"></div>
                    <button type="submit" class="btn btn-primary btn-block btn-lg mb-4" id="submit">Sign in</button>
                </form>
            </div>
            <div class="col-4"></div>
        </div>
    </div>
</template>
