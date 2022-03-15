<template>
  <div class="editor-root">
    <textarea
      v-model.lazy="innerText"
      class="textarea"
      rows="10"
      placeholder="answer"
    />
    <div class="controles">
      <AppButoon class="button" @click="handleSaveClick"> Save </AppButoon>
      <AppButoon class="button" @click="handleCancelClick"> Cancel </AppButoon>
    </div>
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
  cancel: null,
});

const innerText = ref(props.text);

const handleSaveClick = () => {
  emit("save", innerText.value);
};

const handleCancelClick = () => {
  console.log('dddddddddd');
  emit("cancel");
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

.controles {
  display: flex;
  flex-flow: row;
}

.button:not(:last-child) {
  margin-right: 0.25rem;
}
</style>
