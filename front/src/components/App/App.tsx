import type { Component } from 'solid-js';
import ProcessBar from '../ProcessBar';
import Form from '../Form';

import styles from './App.module.css';

export const App: Component = () => {
  const content = "Merit Network, Inc., an independent non-profit 501(c)(3) corporation governed by Michigan''s public universities, was formed in 1966 as the Michigan Educational Research Information Triad to explore computer networking between three of Michigan''s public universities as a means to help the state''s educational and economic development. With initial support from the State of Michigan and the National Science Foundation (NSF), the packet-switched network was first demonstrated in December 1971 when an interactive host to host connection was made between the IBM mainframe computer systems at the University of Michigan in Ann Arbor and Wayne State University in Detroit. In October 1972 connections to the CDC mainframe at Michigan State University in East Lansing completed the triad. Over the next several years in addition to host to host interactive connections the network was enhanced to support terminal to host connections, host to host batch connections (remote job submission, remote printing, batch file transfer), interactive file transfer, gateways to the Tymnet and Telenet public data networks, X.25 host attachments, gateways to X.25 data networks, Ethernet attached hosts, and eventually TCP/IP and additional public universities in Michigan join the network. All of this set the stage for Merit''s role in the NSFNET project starting in the mid-1980s.\n\n".repeat(4);

  return (
    <div class={styles.App}>
      <ProcessBar current={1} total={3} />
      <Form
        content={content}
        question='What completed the triad '
        choices={[
          '7,000,000 square kilometres',
          '7,000,000 square kilometres (2,700,000 sq mi), of which 5,500,000'
        ]}
      />
    </div>
  );
};