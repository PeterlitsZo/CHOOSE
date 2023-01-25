import { Component, createEffect, createResource, createSignal, Show } from 'solid-js';
import axios from 'axios';

import { BASE_URL } from '../../const';
import ProcessBar from '../ProcessBar';
import Form from '../Form';
import { Choice } from '../Form/Form';

import styles from './App.module.css';

export const App: Component = () => {
  let [userName, setUserName] = createSignal('');

  return (
    <div class={styles.App}>
      <Show when={userName() != ''} fallback={<SetUserName setUserName={setUserName} />}>
        <InnerApp
          userName={userName()}
        />
      </Show>
    </div>
  );
};

interface InnerAppProps {
  userName: string;
}

const InnerApp: Component<InnerAppProps> = (props) => {
  let [questions] = createResource(async () => {
    let result = await axios.get(BASE_URL + '/api/v1/questions');
    return result.data.questions as any as Array<{
      id: string,
      context: string,
      question: string,
      choices: Choice[],
    }>;
  })
  let [answer] = createResource(async () => {
    return await axios.get(BASE_URL + '/api/v1/answers', {
      params: { userName: props.userName }
    })
  })
  let [current, setCurrent] = createSignal(0);
  createEffect(() => {
    console.log('[DEBUG]',
      'questions()', questions(), 'answer()', answer(),
      'current()', current())
    if (answer() === undefined || questions() === undefined) return;
    console.log('[DEBUG]',
      'questions()!.length', questions()!.length,
      'current()', current(),
      'questions()!.length <= current()', questions()!.length <= current()
      );
    (answer()!.data.answers as any[]).forEach((ans) => {
      if (ans.question_id === questions()![current()].id) {
        setCurrent(v => v + 1)
        return
      }
    });
  })

  return (
    <Show when={questions() !== undefined} fallback={<>Loading...</>}>
      <Show when={questions()!.length > current()} fallback={<>Already Done.</>}>
        <ProcessBar current={current() + 1} total={questions()!.length} />
        <Form
          content={questions()![current()].context}
          question={questions()![current()].question}
          choices={questions()![current()].choices}
          submitAndNext={async (answer: {answer: string, isHardToAnswer: boolean}) => {
            await axios.post(BASE_URL + '/api/v1/answers', {
              question_id: questions()![current()].id,
              user_name: props.userName,
              answer: answer.answer,
              is_hard_to_answer: answer.isHardToAnswer,
            })
            setCurrent(v => v + 1)
          }}
        />
      </Show>
    </Show>
  )
}

interface SetUserNameProps {
  setUserName: (userName: string) => void,
}

const SetUserName: Component<SetUserNameProps> = (props) => {
  const [userName, setUserName] = createSignal('');

  return (
    <div class={styles.SetUserName}>
      <div>Enter Your User Name</div>
      <input class={styles.UserNameInput} onInput={e => setUserName(e.currentTarget.value)}/>
      <button class={styles.UserNameButton} onClick={() => props.setUserName(userName())}>Go</button>
    </div>
  )
}