import { createApp } from 'vue'
import App from './App.vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import router from './router'
import { createPinia } from 'pinia'
import * as echarts from 'echarts'
import VChart from 'vue-echarts'
import 'echarts/core'
import { PieChart, LineChart, RadarChart } from 'echarts/charts'
import {
  TitleComponent,
  TooltipComponent,
  LegendComponent,
  GridComponent
} from 'echarts/components'
import { CanvasRenderer } from 'echarts/renderers'

// 注册必须的组件
echarts.use([
  TitleComponent,
  TooltipComponent,
  LegendComponent,
  GridComponent,
  PieChart,
  LineChart,
  RadarChart,
  CanvasRenderer
])

const app = createApp(App)
app.use(ElementPlus)
app.use(router)
app.use(createPinia())
app.component('VChart', VChart)
app.mount('#app')
