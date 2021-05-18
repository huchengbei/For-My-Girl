<template>
  <div class="cody" style="background-color: #E8EFF5">
  <div style="margin: 0 auto;width: 70%;">
    <div class="slideshow js-slideshow slideshow--transition-slide" data-swipe="on">
      <p class="sr-only">Slideshow Items</p>

      <ul class="slideshow__content">

        <template v-for="(slide, index) in slides" :key="index">
          <li class="slideshow__item bg js-slideshow__item" :style="'background-image: url(' + slide.image_link + ');'">
            <div class='container max-width-sm'>
              <div class='text-component text-center'>
                <h1 v-if="slide.title">{{ slide.title }}</h1>
                <p v-if="slide.description">{{ slide.description }}</p>
              </div>
            </div>
          </li>
        </template>

      </ul>

      <ul>
        <li class="slideshow__control js-slideshow__control">
          <button class="reset slideshow__btn js-tab-focus">
            <svg class="icon" viewBox="0 0 32 32"><title>Show previous slide</title><path d="M20.768,31.395L10.186,16.581c-0.248-0.348-0.248-0.814,0-1.162L20.768,0.605l1.627,1.162L12.229,16 l10.166,14.232L20.768,31.395z"></path></svg>
          </button>
        </li>

        <li class="slideshow__control js-slideshow__control">
          <button class="reset slideshow__btn js-tab-focus">
            <svg class="icon" viewBox="0 0 32 32"><title>Show next slide</title><path d="M11.232,31.395l-1.627-1.162L19.771,16L9.605,1.768l1.627-1.162l10.582,14.813 c0.248,0.348,0.248,0.814,0,1.162L11.232,31.395z"></path></svg>
          </button>
        </li>
      </ul>
    </div>
    <div class="position-relative z-index-1 overflow-hidden padding-y-sm bg-contrast-lower" style="background-color: #E8EFF5">
<!--      <div class="margin-bottom-lg">-->
<!--        <h1 class="text-center">Vertical Timeline</h1>-->
<!--      </div>-->

      <div class="container max-width-adaptive-lg">
        <div class="v-timeline v-timeline--icons js-v-timeline" data-animation="on">

          <section class="v-timeline__section js-v-timeline__section" v-for="(moment,index) in moments" :key="index">
            <div class="v-timeline__marker bg-primary border border-3 border-bg shadow-xs" aria-hidden="true">
              <img src="img/cd-icon-picture.svg" v-if="moment.moment_type === 'image'"/>
              <img src="img/cd-icon-picture.svg" v-else-if="moment.moment_type === 'album'"/>
              <img src="img/cd-icon-audio.svg" v-else-if="moment.moment_type === 'music'"/>
              <img src="img/cd-icon-movie.svg" v-else-if="moment.moment_type === 'video'"/>
              <img src="img/cd-icon-text.svg" v-else/>
            </div>

            <div class="v-timeline__items-group">
              <div class="v-timeline__item bg padding-md radius-md shadow-xs">
                <div class="v-timeline__date margin-bottom-sm">
                  <time class="v-timeline__date-value" datetime="2021-03-18T00:00">{{ format_date(moment.datetime) }}</time>
                </div>

                <div class="text-component">
                  <h2 v-if="moment.link"><a :href="moment.link" target="_blank">{{ moment.title }}</a></h2>
                  <h2 v-else>{{ moment.title }}</h2>

                  <p>{{ moment.content }}</p>

                  <div class="media">
                    <template v-if="moment.moment_type === 'image'">
                      <figure class="text-component__block">
                        <img :src="moment.image_link" alt="Image description"/>
                      </figure>
                    </template>
                    <template v-else-if="moment.moment_type === 'album'">
                      <figure class="text-component__block">
                        <img :src="moment.album_link" alt="Image description"/>
                      </figure>
                    </template>
                    <template v-else-if="moment.moment_type === 'music'">
                      <figure class="text-component__block">
                        <audio :src="moment.music_link" controls="controls" type="audio/mpeg" alt="Image description"/>
                      </figure>
                    </template>
                    <template v-else-if="moment.moment_type === 'video'">
<!--                      <figure class="text-component__block">-->
                        <video :src="moment.video_link" controls="controls" alt="Image description"/>
<!--                      </figure>-->
                    </template>
                  </div>
                </div>

<!--                <div class="margin-top-md">-->
<!--                  <a class="btn btn&#45;&#45;primary" href="#0">Read More</a>-->
<!--                </div>-->
              </div>
            </div>
          </section>
        </div>
      </div>
    </div>
  </div>
  </div>
</template>

<script>
export default {
  name: 'Cody',
  data() {
    return {
      js_loaded: false,
      slides: [],
      moments: []
    }
  },
  methods: {
    format_date(d) {
      let date = new Date(d)
      let year = date.getFullYear()
      let month = date.getMonth() + 1
      let day = date.getDate()
      return '' + year + '年' + month + '月' + day + '日'
    },
    load_js() {
      let swipe = document.createElement('script');
      swipe.setAttribute('src', './js/swipe.js'); // 👈 make sure to use the correct path
      swipe.setAttribute('id', 'swipe-js');

      let slideshow = document.createElement('script');
      slideshow.setAttribute('src', './js/slideshow.js'); // 👈 make sure to use the correct path
      slideshow.setAttribute('id', 'slideshow-js');

      let vertical_timeline = document.createElement('script');
      vertical_timeline.setAttribute('src', './js/vertical-timeline.js'); // 👈 make sure to use the correct path
      vertical_timeline.setAttribute('id', 'vertical-timeline-js');

      document.body.appendChild(swipe);
      document.body.appendChild(slideshow);
      document.body.appendChild(vertical_timeline);
      this.js_loaded = true
    },
    remove_js(){
      let swipe =document.getElementById('swipe-js')
      document.body.removeChild(swipe)
      let slideshow =document.getElementById('slideshow-js')
      document.body.removeChild(slideshow)
      let vertical_timeline =document.getElementById('vertical-timeline-js')
      document.body.removeChild(vertical_timeline)
      this.js_loaded = false
    }
  },
  created() {
    let that = this
    this.axios.get('/api/slides').then((res) => {
      if (200 !== res.data.code) {
        return
      }
      let data = res.data.data
      that.slides = data
      if (that.js_loaded) {
        that.remove_js()
      }
      that.load_js()
    })
    this.axios.get('/api/moments').then((res) => {
      if (200 !== res.data.code) {
        return
      }
      let data = res.data.data
      that.moments = data
      if (that.js_loaded) {
        that.remove_js()
      }
      that.load_js()
    })
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">


</style>
