import { Component, createSignal, For, Show } from "solid-js";

import styles from './Form.module.css';

interface FormProps {
  content: string;
  question: string;
  choices: string[];
}

type Answer = {
  state: 'unanswered'
  answer: undefined,
  isHardToAnswer: undefined,
} | {
  state: 'unlabeled',
  answer: string,
  isHardToAnswer: undefined,
} | {
  state: 'ready-to-submit',
  answer: string,
  isHardToAnswer: boolean,
}

export const Form: Component<FormProps> = (props) => {
  const [answer, setAnswer] = createSignal({state: 'unanswered'} as Answer)

  const setInnerAnswer = (c: string) => () => {
    const a = answer();
    if (a.state === 'unanswered') {
      setAnswer({ ...a, state: 'unlabeled', answer: c });
    } else {
      setAnswer({ ...a, answer: c })
    }
  };

  return (
    <div class={styles.FormWrapper}>
      <div class={styles.Form}>
        <div class={styles.Content}>{props.content}</div>

        <div class={styles.Question}>
          <div class={styles.Title}>
            <span class={styles.TitlePrefix}>Question: </span>
            {props.question}
          </div>
          <For each={props.choices}>{c => (
            <div
              onClick={setInnerAnswer(c)}
              classList={{
                [styles.Choice]: true,
                [styles.choosed]: c === answer().answer
              }}
            >{c}</div>
          )}</For>
          <div
            onClick={setInnerAnswer("I don't know")}
            classList={{
              [styles.Choice]: true,
              [styles.choosed]: "I don't know" === answer().answer,
              [styles.withoutBottomBorder]: true,
            }}
          >I don't know</div>
        </div>

        <Show when={answer().state != 'unanswered'}>
          <div class={styles.Question}>
            <div class={styles.Title}>
              <span class={styles.TitlePrefix}>Is this question hard to answer? </span>
            </div>
            <div
              onClick={() => setAnswer(a => ({
                state: 'ready-to-submit', answer: a.answer!, isHardToAnswer: true
              }))}
              classList={{
                [styles.Choice]: true,
                [styles.choosed]: answer().isHardToAnswer === true
              }}
            >Yes</div>
            <div
              onClick={() => setAnswer(a => ({
                state: 'ready-to-submit', answer: a.answer!, isHardToAnswer: false
              }))}
              classList={{
                [styles.Choice]: true,
                [styles.withoutBottomBorder]: true,
                [styles.choosed]: answer().isHardToAnswer === false
              }}
            >No</div>
          </div>
        </Show>
      </div>
    </div>
  );
}