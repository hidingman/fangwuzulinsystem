import Vue from 'vue';
//配置路由
import VueRouter from 'vue-router'
Vue.use(VueRouter);
//1.创建组件
import Index from '@/views/index'
import Home from '@/views/home'
import Login from '@/views/login'
import NotFound from '@/views/404'
import UpdatePassword from '@/views/update-password'
import pay from '@/views/pay'
import register from '@/views/register'
import center from '@/views/center'
	import news from '@/views/modules/news/list'
	import yuangong from '@/views/modules/yuangong/list'
	import fangyuanxinxi from '@/views/modules/fangyuanxinxi/list'
	import fangwuhuxing from '@/views/modules/fangwuhuxing/list'
	import fangwuzulin from '@/views/modules/fangwuzulin/list'
	import zufangyixiang from '@/views/modules/zufangyixiang/list'
	import storeup from '@/views/modules/storeup/list'
	import yuyuekanfang from '@/views/modules/yuyuekanfang/list'
	import discussfangyuanxinxi from '@/views/modules/discussfangyuanxinxi/list'
	import zulinhetong from '@/views/modules/zulinhetong/list'
	import chat from '@/views/modules/chat/list'
	import guke from '@/views/modules/guke/list'
	import zulinpingjia from '@/views/modules/zulinpingjia/list'
	import kaoqindaka from '@/views/modules/kaoqindaka/list'
	import tuizushenqing from '@/views/modules/tuizushenqing/list'
	import config from '@/views/modules/config/list'
	import newstype from '@/views/modules/newstype/list'


//2.配置路由   注意：名字
export const routes = [{
	path: '/',
	name: '系统首页',
	component: Index,
	children: [{
		// 这里不设置值，是把main作为默认页面
		path: '/',
		name: '系统首页',
		component: Home,
		meta: {icon:'', title:'center', affix: true}
	}, {
		path: '/updatePassword',
		name: '修改密码',
		component: UpdatePassword,
		meta: {icon:'', title:'updatePassword'}
	}, {
		path: '/pay',
		name: '支付',
		component: pay,
		meta: {icon:'', title:'pay'}
	}, {
		path: '/center',
		name: '个人信息',
		component: center,
		meta: {icon:'', title:'center'}
	}
	,{
		path: '/news',
		name: '公告资讯',
		component: news
	}
	,{
		path: '/yuangong',
		name: '员工',
		component: yuangong
	}
	,{
		path: '/fangyuanxinxi',
		name: '房源信息',
		component: fangyuanxinxi
	}
	,{
		path: '/fangwuhuxing',
		name: '房屋户型',
		component: fangwuhuxing
	}
	,{
		path: '/fangwuzulin',
		name: '房屋租赁',
		component: fangwuzulin
	}
	,{
		path: '/zufangyixiang',
		name: '租房意向',
		component: zufangyixiang
	}
	,{
		path: '/storeup',
		name: '我的收藏',
		component: storeup
	}
	,{
		path: '/yuyuekanfang',
		name: '预约看房',
		component: yuyuekanfang
	}
	,{
		path: '/discussfangyuanxinxi',
		name: '房源信息评论',
		component: discussfangyuanxinxi
	}
	,{
		path: '/zulinhetong',
		name: '租赁合同',
		component: zulinhetong
	}
	,{
		path: '/chat',
		name: '智能客服',
		component: chat
	}
	,{
		path: '/guke',
		name: '顾客',
		component: guke
	}
	,{
		path: '/zulinpingjia',
		name: '租赁评价',
		component: zulinpingjia
	}
	,{
		path: '/kaoqindaka',
		name: '考勤打卡',
		component: kaoqindaka
	}
	,{
		path: '/tuizushenqing',
		name: '退租申请',
		component: tuizushenqing
	}
	,{
		path: '/config',
		name: '轮播图管理',
		component: config
	}
	,{
		path: '/newstype',
		name: '公告资讯分类',
		component: newstype
	}
	]
	},
	{
		path: '/login',
		name: 'login',
		component: Login,
		meta: {icon:'', title:'login'}
	},
	{
		path: '/register',
		name: 'register',
		component: register,
		meta: {icon:'', title:'register'}
	},
	{
		path: '*',
		component: NotFound
	}
]
//3.实例化VueRouter  注意：名字
const router = new VueRouter({
	mode: 'hash',
	/*hash模式改为history*/
	routes // （缩写）相当于 routes: routes
})
const originalPush = VueRouter.prototype.push
//修改原型对象中的push方法
VueRouter.prototype.push = function push(location) {
	return originalPush.call(this, location).catch(err => err)
}
export default router;
