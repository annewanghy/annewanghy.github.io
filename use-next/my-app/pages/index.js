import React from "react";
import Head from "next/head";

const Home = () => (
  <div>
    <Head>
      {/* 
      Scripts with async in Head are executed twice 
      https://github.com/zeit/next.js/issues/9070:
      one workaround is to set async="" */}
      <script src="https://annewanghy.github.io/inline-style.js" async="" />
    </Head>
  </div>
);

export default Home;
