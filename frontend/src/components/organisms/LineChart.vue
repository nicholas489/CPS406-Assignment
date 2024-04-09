<template>
    <div class="line-chart-container">
      <canvas ref="chart" />
    </div>
  </template>
  
  <script setup>
  import { defineComponent, ref, watchEffect } from 'vue';
  import { Line } from 'vue-chartjs';
  import {
    Chart as ChartJS,
    Title,
    Tooltip,
    Legend,
    LineElement,
    CategoryScale,
    LinearScale,
    PointElement,
  } from 'chart.js';
  
  ChartJS.register(Title, Tooltip, Legend, LineElement, CategoryScale, LinearScale, PointElement);
  
defineComponent({
    name: 'LineChart',
    components: {
        Line,
    },
    props: {
        expenses: Array,
        profits: Array,
    },
    setup(props) {
        const chartRef = ref(null);
        const chartData = ref({
            labels: [],
            datasets: [],
        });
        const chartOptions = {
            responsive: true,
            maintainAspectRatio: false,
        };

        watchEffect(() => {
            const dataLength = Math.min(props.expenses.length, props.profits.length, 25);
            chartData.value = {
                labels: [...Array(dataLength).keys()].map(n => `Month ${n + 1}`),
                datasets: [
                    {
                        label: 'Expenses',
                        data: props.expenses.slice(-dataLength),
                        borderColor: 'red',
                        fill: false,
                    },
                    {
                        label: 'Profits',
                        data: props.profits.slice(-dataLength),
                        borderColor: 'green',
                        fill: false,
                    },
                ],
            };
        });

        onMounted(() => {
            new ChartJS(chartRef.value.getContext('2d'), {
                type: 'line',
                data: chartData.value,
                options: chartOptions,
            });
        });

        return {
            chartRef,
            chartData,
            chartOptions,
        };
    },
});
  </script>
  
  <style scoped>
  .line-chart-container {
    position: relative;
    width: 100%;
    height: 100%;
  }
  </style>
  