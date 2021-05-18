<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <div class="demo-box">
      <div>
        <h3>竖直 - 左轴</h3>
        <Timeline line-position="left" direction="vertical" fill-parent fill-node>
          <TimelineItem v-for="i in length" :key="i" :time="time" :title="getTitle(i)" :node="`${i}-${i}`"
                         :content="getContent(i)"/>
        </Timeline>
      </div>
      <div>
        <h3>竖直 - 中轴</h3>
        <Timeline line-style="dotted" line-position="center" direction="vertical"
                  fill-parent node-shape="rect" show-endline>
          <TimelineItem v-for="i in length" :key="i" :time="time" :title="getTitle(i)"
                         :position="i % 2 ? 'left' : 'right'" :content="getContent(i)"/>
        </Timeline>
      </div>
      <div>
        <h3>竖直 - 右轴</h3>
        <Timeline line-style="dashed" line-position="right" fill-parent direction="vertical">
          <TimelineItem v-for="i in length" :key="i" :time="time" :title="getTitle(i)"
                         :content="getContent(i)"/>
        </Timeline>
      </div>
    </div>
    <h3>heeeeeee</h3>
  </div>
</template>

<script>
// import {Timeline, TimelineItem} from '@hyjiacan/vue-timeline'

export default {
  name: 'HelloWorld',
  data() {
    return {
      length: 3,
      time: new Date().toLocaleString(),
      colors: [
        '',
        'orange',
        'red',
        'green',
        'lightblue'
      ],
    }
  },
  components: {
  },
  methods: {
    getTitle(v) {
      return 'tttttttttttttttt' + v
    },
    getContent(v) {
      return 'cccccccccccc' + v
    }
  },
  props: {
    msg: String
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.demo-box {
  /*display: grid;*/
  /*grid-template-columns: 25% auto 25%;*/

  > div {
    border-top: 1px solid #888888;
  }
}

.custom-node {
  display: block;
  box-sizing: border-box;
}

.custom-content {
  padding: 10px 20px;
}

.custom-node--content{
  display: inline-block;
  width: 18px;
  height: 18px;
}
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
.demo-box {
  /*display: grid;*/
  /*grid-template-columns: 25% auto 25%;*/

  > div {
    border-top: 1px solid #888888;
  }
}

.custom-node {
  display: block;
  box-sizing: border-box;
}

.custom-content {
  padding: 10px 20px;
}

.custom-node--content{
  display: inline-block;
  width: 18px;
  height: 18px;
}
</style>
