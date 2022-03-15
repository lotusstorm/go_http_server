<template>
  <div class="container">
    <textarea v-model="question" class="textarea" rows="10" placeholder="question" />
    <div class="controles">
      <label>
        switch
        <input v-model="isMulty" type="checkbox" />
      </label>
      <AppButton @click="handleIsEditClick"> + </AppButton>
    </div>

    <AppEditor
      v-if="isEdit"
      @save="handleAddClick"
      @cancel="handleCancelClick"
    />
    <div class="answers-list">
      <div v-for="answer of answersList" :key="answer.id" class="answers-item">
        <AppAnswerItem
          :index="answer.id"
          :text="answer.text"
          @save="handleUpdateClick($event, answer)"
        >
          <template #prepend>
            <AppAnswerPicker
              v-model="selectedAnswers"
              :value="answer.id"
              :is-multy="isMulty"
              class="answer-picker"
            />
          </template>
        </AppAnswerItem>
      </div>
    </div>
    <AppButton @click="handlSeaveClick"> Save </AppButton>
  </div>
</template>

<script setup>
import { reactive, ref, watch } from "vue";
import AppAnswerItem from "@/components/AppAnswerItem.vue";
import AppButton from "./components/AppButton.vue";
import AppEditor from "@/components/AppEditor.vue";
import AppAnswerPicker from "./components/AppAnswerPicker.vue";

const answersList = reactive([]);
const selectedAnswers = ref([]);
const question = ref("");

const isEdit = ref(false);
const isMulty = ref(false);

watch(isMulty, () => {
  selectedAnswers.value = [];
});

const handleIsEditClick = () => {
  isEdit.value = true;
};

const handleAddClick = (value) => {
  isEdit.value = false;

  answersList.push({
    id: answersList.length,
    text: value,
  });
};

const handleUpdateClick = (text, item) => {
  item.text = text;
};

const handleCancelClick = () => {
  isEdit.value = false;
};

const handlSeaveClick = () => {
  // const quiz = {
  //   question: question.value,
  //   answers: selectedAnswers.value,
  //   answer: answersList.value,
  //   multySelect: isMulty.value,
  // };
};
</script>

<style>
@import "./assets/base.css";

#app {
  max-width: 1024px;
  height: 100vh;
  margin: 0 auto;

  font-weight: normal;
  overflow-y: auto;
}

.answer-picker {
  margin-right: 0.25rem;
}

.container {
  display: flex;
  flex-direction: column;
  width: 100%;
  border-radius: 16px;
  border: 1px solid silver;
  padding: 1rem 1rem 0.5rem 1rem;
}

.container > .textarea {
  width: 100%;
  border-radius: 0.5rem;
  padding: 0.35rem;
  margin-bottom: 0.5rem;
}

.answers-list {
}

.answers-item:not(:last-child) {
  margin-bottom: 0.35rem;
}
</style>
