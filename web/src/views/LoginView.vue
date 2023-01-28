<template>
    <main class="container">
        <article class="grid">
            <div>
                <hgroup>
                    <h1>Sign in</h1>
                    <h2>Please sign in to access setting page</h2>
                </hgroup>
                <form>
                    <input type="text" name="username" placeholder="Login" aria-label="Login" autocomplete="nickname" required>
                    <input type="password" name="password" placeholder="Password" aria-label="Password" autocomplete="current-password" required>
                    <button type="submit" class="contrast" onclick="event.preventDefault()">Login</button>
                </form>
            </div>
        </article> 
    </main>
</template>