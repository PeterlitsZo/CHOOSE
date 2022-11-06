import type { Component } from "solid-js";

import styles from './ProcessBar.module.css';

interface ProcessBarProps {
  current: number;
  total: number;
}

export const ProcessBar: Component<ProcessBarProps> = (props) => {
  return (
    <div class={styles.Main}>
      <div class={styles.Sub} style={{
        right: (1 - props.current / props.total) * 100 + '%'
      }} />
    </div>
  )
}