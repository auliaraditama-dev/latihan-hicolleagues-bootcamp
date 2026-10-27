<script setup>
import { computed } from "vue";

const props = defineProps({
  sisaTugas: {
    type: Number,
    required: true,
  },
  totalTugas: {
    type: Number,
    required: true,
  },
});
const selesai = computed(() => {
  return props.totalTugas - props.sisaTugas;
});
const persentase = computed(() => {
  if (props.totalTugas === 0) return 0;
  return Math.round((selesai.value / props.totalTugas) * 100);
});
</script>

<template>
  <div class="stats-wrapper">
    <div class="stats-card">
      <span
        >Total: <strong>{{ totalTugas }}</strong></span
      >
      <span
        >Selesai: <strong>{{ selesai }}</strong></span
      >
      <span
        >Sisa: <strong class="badge">{{ sisaTugas }}</strong></span
      >
    </div>

    <div class="progress-section" v-if="totalTugas > 0">
      <div class="progress-label">
        <span>Progress:</span>
        <span class="progress-value">{{ persentase }}%</span>
      </div>

      <div class="progress-track">
        <div
          class="progress-fill"
          :class="{ 'is-completed': persentase === 100 }"
          :style="{ width: persentase + '%' }"
        ></div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.stats-wrapper {
  margin-top: 20px;
}

.stats-card {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  background: #f3e8ff;
  border-radius: 8px;
  font-size: 14px;
  color: #5b21b6;
}

.badge {
  background: #7c3aed;
  color: white;
  padding: 2px 8px;
  border-radius: 9999px;
  font-size: 12px;
}

.progress-section {
  margin-top: 12px;
}

.progress-label {
  display: flex;
  justify-content: space-between;
  font-size: 13px;
  color: #6b5b86;
  margin-bottom: 6px;
  font-weight: 500;
}

.progress-track {
  width: 100%;
  height: 10px;
  background-color: #ede9fe;
  border-radius: 9999px;
  overflow: hidden;
}

.progress-fill {
  height: 100%;
  background-color: #7c3aed;
  border-radius: 9999px;
  transition:
    width 0.3s ease,
    background-color 0.3s ease;
}

.progress-fill.is-completed {
  background-color: #10b981;
}
</style>
