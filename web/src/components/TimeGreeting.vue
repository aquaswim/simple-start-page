<script lang="ts">
export default {
  data() {
    return {
      greetingMsg: "Hello World!",
    };
  },
  created(this) {
    var hour = new Date().getHours();
    var dayParts = [24, 18, 12]; //Keep in this order
    var greetings = ["Good Evening", "Good Afternoon", "Good Morning"];
    this.greetingMsg = "Hello World!";
    for (const [i, dayPart] of dayParts.entries()) {
      if (hour < dayPart) {
        this.greetingMsg = greetings[i];
      }
    }
  },
};
</script>

<template>
  <h1>
    <u>{{ greetingMsg }}</u>
  </h1>
</template>
