<template>
  <div class="editor-root">
    <textarea
      v-model.lazy="innerText"
      class="textarea"
      rows="10"
      placeholder="answer"
    />
    <AppButoon class="button" @click="handleSaveClick"> Save </AppButoon>
  </div>
</template>

<script setup>
import AppButoon from "@/components/AppButton.vue";
import { ref } from "vue";

const props = defineProps({
  text: {
    type: String,
    default: "",
  },
});

const emit = defineEmits({
  save: null,
});

const innerText = ref(props.text);

const handleSaveClick = () => {
  emit("save", innerText.value);
};
</script>

<style scoped>
.editor-root {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  width: 100%;
}

.textarea {
  width: 100%;
  border-radius: 0.5rem;
  padding: 0.35rem;
  margin-bottom: 0.5rem;
}
</style>
