import React from "react";
import Head from "next/head";

const Ext2 = () => <script src="/static/ext2.js" async></script>;

const Home = () => (
  <div>
    <Head>
      {/* refer:https://github.com/zeit/next.js/issues/9070 */}
      <script src="/static/ext.js"></script>
      {/* Use component instead of direct set script in Head to fixexecute twice */}
      <Ext2></Ext2>
    </Head>
  </div>
);

export default Home;
