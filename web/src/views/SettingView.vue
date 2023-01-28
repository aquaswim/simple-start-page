<template>
    <main class="container">
        <form>
            <label for="input-site-title">Text</label>
            <input type="text" id="input-site-title" name="site_title" placeholder="My Site">
            
            <label for="input-about">About page content</label>
            <textarea name="about" id="input-about" rows="10" placeholder="This is my simple site"></textarea>
            
            <fieldset>
                <label for="input-hide-setting">
                    <input type="checkbox" id="input-hide-setting" name="hide_setting" role="switch">
                    Hide setting page
                </label>
            </fieldset>
            
            <input type="submit" value="Save">
        </form>
    </main>
</template>