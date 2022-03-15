<template>
  <div class="answer-item-root">
    <AppEditor
      v-if="isEdit"
      :text="text"
      @save="handleSaveClick"
      @cancel="handleCancelClick"
    />
    <div v-else class="preview">
      <div class="item">
        <slot name="prepend" />
        <span class="index">{{ itemIndex }} )</span>
        <p class="content">{{ text }}</p>

        <div class="controles">
          <AppButoon class="button" @click="handleIsEditClick"> E </AppButoon>
          <AppButoon class="button"> D </AppButoon>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, ref } from "vue";
import AppButoon from "@/components/AppButton.vue";
import AppEditor from "@/components/AppEditor.vue";

const props = defineProps({
  text: {
    type: String,
    default: "",
  },
  index: {
    type: Number,
    required: true,
  },
});

const emit = defineEmits({
  save: null,
});

const isEdit = ref(false);

const itemIndex = computed(() => String.fromCharCode(97 + props.index));

const handleIsEditClick = () => {
  isEdit.value = true;
};

const handleSaveClick = (value) => {
  isEdit.value = false;
  emit("save", value);
};

const handleCancelClick = () => {
  isEdit.value = false;
};
</script>

<style scoped>
.preview {
  --controles-visibility: hidden;

  display: flex;
  flex-flow: row;
  align-items: center;
}

.preview:hover {
  --controles-visibility: visible;
}

.item {
  display: flex;
  flex-flow: row;
  position: relative;
  width: 100%;
  padding: 0.4rem;
  border-radius: 0.25rem;
  border: 1px solid silver;
  align-items: center;
}

.index {
  font-weight: 600;
  margin-right: 0.25rem;
}
.content {
}

.controles {
  display: flex;
  flex-flow: row;
  position: absolute;
  visibility: var(--controles-visibility);
  right: 0;
  top: 0;
  bottom: 0;
  align-items: center;
}

.button:not(:last-child) {
  margin-right: 0.25rem;
}
</style>
