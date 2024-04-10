<template>
    <div class="box">
        <h1>Coach Dashboard</h1>
        <h3>Welcome, "name"!</h3>
    </div>
</template>

<script setup>
import LineChart from '../organisms/LineChart.vue';

const expenses = [1200, 1400, 800, 1600, 1300, 950, 1200, 1100, 1450, 950, 1250, 1450, 1050, 950, 1200, 1150, 1250, 1350, 1450, 1550, 1650, 1750, 1850, 1950, 2050];
const profits = [1500, 1600, 1700, 1800, 1900, 2000, 1800, 1700, 1600, 1900, 2200, 2000, 2100, 2000, 1900, 1800, 1700, 1600, 1800, 1900, 2100, 2300, 2400, 2200, 2100];
</script>


<style scoped>
.box {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    border: 2px solid #ccc;
    border-radius: 8px;
    height: 300px;
    padding: 20px;
}
</style>
  