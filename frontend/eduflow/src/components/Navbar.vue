<script setup>
    const authed = (localStorage.getItem("token") != null)
</script>

<template>
    <nav class="navbar navbar-expand-lg bg-body-tertiary">
        <div class="container">
            <a class="navbar-brand" href="/">Eduflow</a>
            <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
            </button>
            <div class="collapse navbar-collapse" id="navbarSupportedContent">
                <ul class="navbar-nav me-auto mb-2 mb-lg-0">
                    <li class="nav-item">
                        <!-- <a class="nav-link active" href="/">Home</a> -->
                        <router-link class="nav-link" to="/" exact-active-class="active" active-class="active">Home</router-link>
                    </li>
                    <li class="nav-item">
                        <!-- <a class="nav-link" href="/groups/">Groups</a> -->
                        <router-link class="nav-link" to="/groups/" exact-active-class="active" active-class="active">Groups</router-link>
                    </li>
                    <li class="nav-item">
                        <!-- <a class="nav-link" href="/summaries/">Summaries</a> -->
                        <router-link class="nav-link" to="/summaries/" exact-active-class="active" active-class="active">Summaries</router-link>
                    </li>
                    <!-- <li class="nav-item dropdown">
                        <a class="nav-link dropdown-toggle" href="#" role="button" data-bs-toggle="dropdown" aria-expanded="false">
                            Dropdown
                        </a>
                        <ul class="dropdown-menu">
                            <li><a class="dropdown-item" href="#">Action</a></li>
                            <li><a class="dropdown-item" href="#">Another action</a></li>
                            <li><hr class="dropdown-divider"></li>
                            <li><a class="dropdown-item" href="#">Something else here</a></li>
                        </ul>
                    </li> -->
                </ul>
                <form v-if="!authed" class="d-flex" role="auth">
                    <a class="sign-in btn" href="/auth/sign-in/">Sign in</a>
                    <a class="sign-up btn" href="/auth/sign-up">Sign up</a>
                </form>
                <form v-else class="d-flex" role="profile">
                    <a href="/profile" class="profile btn">Profile</a>
                </form>
            </div>
        </div>
    </nav>
</template>
