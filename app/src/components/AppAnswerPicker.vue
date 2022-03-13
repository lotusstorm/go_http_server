<template>
  <input v-if="isMulty" v-model="innerValue" :value="value" type="checkbox" />
  <input v-else v-model="innerValue" :value="value" type="radio" />
</template>

<script setup>
import { computed } from "vue";

const props = defineProps({
  isMulty: {
    type: Boolean,
    default: false,
  },
  modelValue: {
    type: Array,
    default: () => [],
  },
  value: {
    type: Number,
    required: true,
  },
});

const emit = defineEmits({
  "update:modelValue": null,
});

const innerValue = computed({
  get() {
    if (props.isMulty) {
      return props.modelValue;
    }

    return props.modelValue[0];
  },
  set(val) {
    emit("update:modelValue", props.isMulty ? val : [val]);
  },
});
</script>
