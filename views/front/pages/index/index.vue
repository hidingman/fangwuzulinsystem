<template>
	<view class="content">
		<view :style='{"width":"100%","padding":"0 0 40rpx 0","background":"#FFFBF3","flexDirection":"column","display":"flex","height":"auto"}'>
			<view class="list-swiper-4" :style='{"width":"100%","position":"relative","height":"360rpx"}' @touchstart="touchStart" @touchmove="touchMove" @touchend="touchEnd">
				<view :style='{"width":"100%","position":"absolute","top":"0%","left":"0%","background":"#fff","height":"360rpx"}' class="item animate__animated" :class="prevNumList4 == index  ? 'animate__backOutRight' : (numList4 == index  ? 'animate__backInLeft' : '')" v-for="(swiper,index) in swiperList" :key="index" v-if="numList4 == index || prevNumList4 == index">
					<image :style='{"width":"100%","objectFit":"cover","display":"block","height":"360rpx"}' mode="aspectFill" :src="baseUrl+swiper.img" @tap="onSwiperTap(swiper)"></image>
					<view :style='{"padding":"8rpx 20rpx","margin":"-20rpx 0 0 0","transform":"translate3d(-50%, -50%, 0)","top":"50%","color":"#000","left":"50%","background":"rgba(255, 255, 255, 0.3)","display":"none","lineHeight":"1.5","fontSize":"40rpx","position":"absolute"}'>{{ swiper.title }}</view>
				</view>
				<view class="animate__navigation" :style='{"alignItems":"center","left":"0%","bottom":"0%","background":"rgba(0,0,0,.3)","display":"flex","width":"100%","position":"absolute","justifyContent":"center","height":"40rpx"}'>
					<block v-for="(swiper,index) in swiperList" :key="index">
						<text class="navigation-item" v-if="numList4 == index" :style='{"width":"16rpx","margin":"0 4rpx","borderRadius":"100%","background":"#F7EFDB","height":"16rpx"}'></text>
						<text class="navigation-item" v-if="numList4 != index" :style='{"width":"16rpx","margin":"0 4rpx","borderRadius":"100%","background":"#fff","height":"16rpx"}'></text>
					</block>
				</view>
			</view>
			<!-- menu -->
			<view v-if="true" class="menu_view">
				<block v-for="(item,index1) in menuList" v-bind:key="item.roleName">
					<block v-if="index1==0" v-bind:key="index" v-for=" (menu,index) in item.frontMenu">
						<block v-bind:key="sort" v-for=" (child,sort) in menu.child">
							<block v-bind:key="sort2" v-for=" (button,sort2) in child.buttons">
								<view class="menu-item" v-if="button=='查看' && child.tableName!='yifahuodingdan' && child.tableName!='yituikuandingdan' &&child.tableName!='yiquxiaodingdan' && child.tableName!='weizhifudingdan' && child.tableName!='yizhifudingdan' && child.tableName!='yiwanchengdingdan' " @tap="onPageTap2(child.tableName)">
									<view class="iconarr" :class="child.appFrontIcon"></view>
									
									<view class="text">{{child.menu.split("列表")[0]}}</view>
								</view>
							</block>
						</block>
					</block>
				</block>
			</view>
			<!-- 商品推荐 -->
			<view class="listBox recommend" :style='{"padding":"0 20rpx","margin":"40rpx 0 0","background":"linear-gradient( 135deg, #E4EFBE 0%, #EFE0BE 100%)","order":"1"}'>
				<view class="title" :style='{"padding":"0","margin":"0","background":"url(http://codegen.caihongy.cn/20250210/f6078ca119e2499f82ca68a92f6a1baa.png) no-repeat center / 100% 100%","display":"flex","width":"100%","position":"relative","justifyContent":"center","height":"80rpx"}'>
					<view :style='{"fontSize":"32rpx","lineHeight":"60rpx","color":"#FFFFFF","fontWeight":"500"}'>房源信息推荐</view>
				</view>
				<!-- 样式8 -->
				<view class="list-box style8" :style='{"width":"100%","padding":"24rpx 0","margin":"0","height":"auto"}'>
					<view v-if="fangyuanxinxilist.length > 0" @tap="onDetailTap('fangyuanxinxi',fangyuanxinxilist[0].id)" class="box box1"  :style='{"width":"100%","padding":"0","margin":"0 0 20rpx","position":"relative","height":"auto"}'>
						<image :style='{"width":"100%","objectFit":"cover","display":"block","height":"270rpx"}' class="list-item-image" mode="aspectFill" v-if="fangyuanxinxilist[0].fangwutupian.substring(0,4)=='http'" :src="fangyuanxinxilist[0].fangwutupian"></image>
						<image :style='{"width":"100%","objectFit":"cover","display":"block","height":"270rpx"}' class="list-item-image" mode="aspectFill" v-else :src="fangyuanxinxilist[0].fangwutupian?baseUrl+fangyuanxinxilist[0].fangwutupian.split(',')[0]:''"></image>
						<view :style='{"width":"100%","padding":"8rpx 12rpx","background":"rgba(0,0,0,.3)"}'>
							<view :style='{"padding":"0 0px","overflow":"hidden","whiteSpace":"nowrap","color":"#fff","lineHeight":"1.5","fontSize":"28rpx","textOverflow":"ellipsis"}' class="title ">{{fangyuanxinxilist[0].fangwumingcheng}}</view>
							<view :style='{"padding":"0 0px","overflow":"hidden","whiteSpace":"nowrap","color":"#fff","lineHeight":"1.5","fontSize":"28rpx","textOverflow":"ellipsis"}' class="title ">{{fangyuanxinxilist[0].fangwuhuxing}}</view>
							<view :style='{"padding":"0 0px"}'>
								<text class="icon iconfont icon-shijian21" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#fff"}'></text>
								<text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"24rpx"}'>{{fangyuanxinxilist[0].addtime.split(' ')[0].replace(/\-/g,'-')}}</text>
							</view>
							<view :style='{"padding":"0 0px"}'>
								<text class="icon iconfont icon-geren16" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#fff"}'></text>
								<text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"24rpx"}'>{{fangyuanxinxilist[0].gonghao}}</text>
							</view>
							<view :style='{"padding":"0 0px","display":"inline-block"}'>
								<text class="icon iconfont icon-zan10" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#fff"}'></text>
								<text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"24rpx"}'>{{fangyuanxinxilist[0].thumbsupnum}}</text>
							</view>
							<view :style='{"padding":"0 0px","display":"inline-block"}'>
								<text class="icon iconfont icon-shoucang10" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#fff"}'></text>
								<text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"24rpx"}'>{{fangyuanxinxilist[0].storeupnum}}</text>
							</view>
							<view :style='{"padding":"0 0px","display":"inline-block"}'>
								<text class="icon iconfont icon-chakan2" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#fff"}'></text>
								<text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"24rpx"}'>{{fangyuanxinxilist[0].clicknum}}</text>
							</view>
						</view>
					</view>
					<view class="list-item" :style='{"width":"100%","padding":"0","margin":"0 0 20rpx","justifyContent":"space-between","display":"flex","height":"auto"}'>
						<view v-if="fangyuanxinxilist.length > 1" @tap="onDetailTap('fangyuanxinxi',fangyuanxinxilist[1].id)" class="box box2" :style='{"width":"32%","padding":"0","margin":"0","position":"relative","height":"auto"}'>
							<image :style='{"width":"100%","objectFit":"cover","display":"block","height":"240rpx"}' class="list-item-image" mode="aspectFill" v-if="fangyuanxinxilist[1].fangwutupian.substring(0,4)=='http'" :src="fangyuanxinxilist[1].fangwutupian"></image>
							<image :style='{"width":"100%","objectFit":"cover","display":"block","height":"240rpx"}' class="list-item-image" mode="aspectFill" v-else :src="fangyuanxinxilist[1].fangwutupian?baseUrl+fangyuanxinxilist[1].fangwutupian.split(',')[0]:''"></image>
							<view :style='{"width":"100%","padding":"8rpx 12rpx","background":"rgba(0,0,0,.3)"}'>
								<view :style='{"padding":"0 0px","overflow":"hidden","whiteSpace":"nowrap","color":"#fff","lineHeight":"1.5","fontSize":"28rpx","textOverflow":"ellipsis"}' class="title ">{{fangyuanxinxilist[1].fangwumingcheng}}</view>
								<view :style='{"padding":"0 0px","overflow":"hidden","whiteSpace":"nowrap","color":"#fff","lineHeight":"1.5","fontSize":"28rpx","textOverflow":"ellipsis"}' class="title ">{{fangyuanxinxilist[1].fangwuhuxing}}</view>
								<view :style='{"padding":"0 0px"}'>
									<text class="icon iconfont icon-shijian21" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#fff"}'></text>
									<text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"24rpx"}'>{{fangyuanxinxilist[1].addtime.split(' ')[0].replace(/\-/g,'-')}}</text>
								</view>
								<view :style='{"padding":"0 0px"}'>
									<text class="icon iconfont icon-geren16" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#fff"}'></text>
									<text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"24rpx"}'>{{fangyuanxinxilist[1].gonghao}}</text>
								</view>
								<view :style='{"padding":"0 0px","display":"inline-block"}'>
									<text class="icon iconfont icon-zan10" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#fff"}'></text>
									<text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"24rpx"}'>{{fangyuanxinxilist[1].thumbsupnum}}</text>
								</view>
								<view :style='{"padding":"0 8rpx 0 0","display":"inline-block"}'>
									<text class="icon iconfont icon-shoucang10" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#fff"}'></text>
									<text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"24rpx"}'>{{fangyuanxinxilist[1].storeupnum}}</text>
								</view>
								<view :style='{"padding":"0 0px","display":"inline-block"}'>
									<text class="icon iconfont icon-chakan2" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#fff"}'></text>
									<text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"24rpx"}'>{{fangyuanxinxilist[1].clicknum}}</text>
								</view>
							</view>
						</view>
						<view v-if="fangyuanxinxilist.length > 2" @tap="onDetailTap('fangyuanxinxi',fangyuanxinxilist[2].id)" class="box box3" :style='{"width":"32%","padding":"0","margin":"0","position":"relative","height":"auto"}'>
							<image :style='{"width":"100%","objectFit":"cover","display":"block","height":"240rpx"}' class="list-item-image" mode="aspectFill" v-if="fangyuanxinxilist[2].fangwutupian.substring(0,4)=='http'" :src="fangyuanxinxilist[2].fangwutupian"></image>
							<image :style='{"width":"100%","objectFit":"cover","display":"block","height":"240rpx"}' class="list-item-image" mode="aspectFill" v-else :src="fangyuanxinxilist[2].fangwutupian?baseUrl+fangyuanxinxilist[2].fangwutupian.split(',')[0]:''"></image>
							<view :style='{"width":"100%","padding":"8rpx 12rpx","background":"rgba(0,0,0,.3)"}'>
								<view :style='{"padding":"0 0px","overflow":"hidden","whiteSpace":"nowrap","color":"#fff","lineHeight":"1.5","fontSize":"28rpx","textOverflow":"ellipsis"}' class="title ">{{fangyuanxinxilist[2].fangwumingcheng}}</view>
								<view :style='{"padding":"0 0px","overflow":"hidden","whiteSpace":"nowrap","color":"#fff","lineHeight":"1.5","fontSize":"28rpx","textOverflow":"ellipsis"}' class="title ">{{fangyuanxinxilist[2].fangwuhuxing}}</view>
								<view :style='{"padding":"0 0px"}'>
									<text class="icon iconfont icon-shijian21" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#fff"}'></text>
									<text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"24rpx"}'>{{fangyuanxinxilist[2].addtime.split(' ')[0].replace(/\-/g,'-')}}</text>
								</view>
								<view :style='{"padding":"0 0px"}'>
									<text class="icon iconfont icon-geren16" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#fff"}'></text>
									<text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"24rpx"}'>{{fangyuanxinxilist[2].gonghao}}</text>
								</view>
								<view :style='{"padding":"0 0px","display":"inline-block"}'>
									<text class="icon iconfont icon-zan10" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#fff"}'></text>
									<text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"24rpx"}'>{{fangyuanxinxilist[2].thumbsupnum}}</text>
								</view>
								<view :style='{"padding":"0 8rpx 0 0","display":"inline-block"}'>
									<text class="icon iconfont icon-shoucang10" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#fff"}'></text>
									<text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"24rpx"}'>{{fangyuanxinxilist[2].storeupnum}}</text>
								</view>
								<view :style='{"padding":"0 0px","display":"inline-block"}'>
									<text class="icon iconfont icon-chakan2" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#fff"}'></text>
									<text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"24rpx"}'>{{fangyuanxinxilist[2].clicknum}}</text>
								</view>
							</view>
						</view>
						<view v-if="fangyuanxinxilist.length > 3" @tap="onDetailTap('fangyuanxinxi',fangyuanxinxilist[3].id)" class="box box4" :style='{"width":"32%","padding":"0","margin":"0","position":"relative","height":"auto"}'>
							<image :style='{"width":"100%","objectFit":"cover","display":"block","height":"240rpx"}' class="list-item-image" mode="aspectFill" v-if="fangyuanxinxilist[3].fangwutupian.substring(0,4)=='http'" :src="fangyuanxinxilist[3].fangwutupian"></image>
							<image :style='{"width":"100%","objectFit":"cover","display":"block","height":"240rpx"}' class="list-item-image" mode="aspectFill" v-else :src="fangyuanxinxilist[3].fangwutupian?baseUrl+fangyuanxinxilist[3].fangwutupian.split(',')[0]:''"></image>
							<view :style='{"width":"100%","padding":"8rpx 12rpx","background":"rgba(0,0,0,.3)"}'>
								<view :style='{"padding":"0 0px","overflow":"hidden","whiteSpace":"nowrap","color":"#fff","lineHeight":"1.5","fontSize":"28rpx","textOverflow":"ellipsis"}' class="title ">{{fangyuanxinxilist[3].fangwumingcheng}}</view>
								<view :style='{"padding":"0 0px","overflow":"hidden","whiteSpace":"nowrap","color":"#fff","lineHeight":"1.5","fontSize":"28rpx","textOverflow":"ellipsis"}' class="title ">{{fangyuanxinxilist[3].fangwuhuxing}}</view>
								<view :style='{"padding":"0 0px"}'>
									<text class="icon iconfont icon-shijian21" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#fff"}'></text>
									<text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"24rpx"}'>{{fangyuanxinxilist[3].addtime.split(' ')[0].replace(/\-/g,'-')}}</text>
								</view>
								<view :style='{"padding":"0 0px"}'>
									<text class="icon iconfont icon-geren16" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#fff"}'></text>
									<text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"24rpx"}'>{{fangyuanxinxilist[3].gonghao}}</text>
								</view>
								<view :style='{"padding":"0 0px","display":"inline-block"}'>
									<text class="icon iconfont icon-zan10" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#fff"}'></text>
									<text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"24rpx"}'>{{fangyuanxinxilist[3].thumbsupnum}}</text>
								</view>
								<view :style='{"padding":"0 8rpx 0 0","display":"inline-block"}'>
									<text class="icon iconfont icon-shoucang10" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#fff"}'></text>
									<text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"24rpx"}'>{{fangyuanxinxilist[3].storeupnum}}</text>
								</view>
								<view :style='{"padding":"0 0px","display":"inline-block"}'>
									<text class="icon iconfont icon-chakan2" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#fff"}'></text>
									<text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"24rpx"}'>{{fangyuanxinxilist[3].clicknum}}</text>
								</view>
							</view>
						</view>
					</view>
					<view class="list-item" :style='{"width":"100%","padding":"0","margin":"0","justifyContent":"space-between","display":"flex","height":"auto"}'>
						<view v-if="fangyuanxinxilist.length > 4" @tap="onDetailTap('fangyuanxinxi',fangyuanxinxilist[4].id)" class="box box5" :style='{"width":"48%","padding":"0","margin":"0","position":"relative","height":"auto"}'>
							<image :style='{"width":"100%","objectFit":"cover","display":"block","height":"328rpx"}' class="list-item-image" mode="aspectFill" v-if="fangyuanxinxilist[4].fangwutupian.substring(0,4)=='http'" :src="fangyuanxinxilist[4].fangwutupian"></image>
							<image :style='{"width":"100%","objectFit":"cover","display":"block","height":"328rpx"}' class="list-item-image" mode="aspectFill" v-else :src="fangyuanxinxilist[4].fangwutupian?baseUrl+fangyuanxinxilist[4].fangwutupian.split(',')[0]:''"></image>
							<view :style='{"width":"100%","padding":"8rpx 12rpx","background":"rgba(0,0,0,.3)"}'>
								<view :style='{"padding":"0 0px","overflow":"hidden","whiteSpace":"nowrap","color":"#fff","lineHeight":"1.5","fontSize":"28rpx","textOverflow":"ellipsis"}' class="title ">{{fangyuanxinxilist[4].fangwumingcheng}}</view>
								<view :style='{"padding":"0 0px","overflow":"hidden","whiteSpace":"nowrap","color":"#fff","lineHeight":"1.5","fontSize":"28rpx","textOverflow":"ellipsis"}' class="title ">{{fangyuanxinxilist[4].fangwuhuxing}}</view>
								<view :style='{"padding":"0 0px"}'>
									<text class="icon iconfont icon-shijian21" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#fff"}'></text>
									<text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"24rpx"}'>{{fangyuanxinxilist[4].addtime.split(' ')[0].replace(/\-/g,'-')}}</text>
								</view>
								<view :style='{"padding":"0 0px"}'>
									<text class="icon iconfont icon-geren16" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#fff"}'></text>
									<text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"24rpx"}'>{{fangyuanxinxilist[4].gonghao}}</text>
								</view>
								<view :style='{"padding":"0 0px","display":"inline-block"}'>
									<text class="icon iconfont icon-zan10" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#fff"}'></text>
									<text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"24rpx"}'>{{fangyuanxinxilist[4].thumbsupnum}}</text>
								</view>
								<view :style='{"padding":"0 20rpx","display":"inline-block"}'>
									<text class="icon iconfont icon-shoucang10" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#fff"}'></text>
									<text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"24rpx"}'>{{fangyuanxinxilist[4].storeupnum}}</text>
								</view>
								<view :style='{"padding":"0 0px","display":"inline-block"}'>
									<text class="icon iconfont icon-chakan2" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#fff"}'></text>
									<text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"24rpx"}'>{{fangyuanxinxilist[4].clicknum}}</text>
								</view>
							</view>
						</view>
						<view v-if="fangyuanxinxilist.length > 5" @tap="onDetailTap('fangyuanxinxi',fangyuanxinxilist[5].id)" class="box box6" :style='{"width":"48%","padding":"0","margin":"0","position":"relative","height":"auto"}'>
							<image :style='{"width":"100%","objectFit":"cover","display":"block","height":"328rpx"}' class="list-item-image" mode="aspectFill" v-if="fangyuanxinxilist[5].fangwutupian.substring(0,4)=='http'" :src="fangyuanxinxilist[5].fangwutupian"></image>
							<image :style='{"width":"100%","objectFit":"cover","display":"block","height":"328rpx"}' class="list-item-image" mode="aspectFill" v-else :src="fangyuanxinxilist[5].fangwutupian?baseUrl+fangyuanxinxilist[5].fangwutupian.split(',')[0]:''"></image>
							<view :style='{"width":"100%","padding":"8rpx 12rpx","background":"rgba(0,0,0,.3)"}'>
								<view :style='{"padding":"0 0px","overflow":"hidden","whiteSpace":"nowrap","color":"#fff","lineHeight":"1.5","fontSize":"28rpx","textOverflow":"ellipsis"}' class="title ">{{fangyuanxinxilist[5].fangwumingcheng}}</view>
								<view :style='{"padding":"0 0px","overflow":"hidden","whiteSpace":"nowrap","color":"#fff","lineHeight":"1.5","fontSize":"28rpx","textOverflow":"ellipsis"}' class="title ">{{fangyuanxinxilist[5].fangwuhuxing}}</view>
								<view :style='{"padding":"0 0px"}'>
									<text class="icon iconfont icon-shijian21" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#fff"}'></text>
									<text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"24rpx"}'>{{fangyuanxinxilist[5].addtime.split(' ')[0].replace(/\-/g,'-')}}</text>
								</view>
								<view :style='{"padding":"0 0px"}'>
									<text class="icon iconfont icon-geren16" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#fff"}'></text>
									<text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"24rpx"}'>{{fangyuanxinxilist[5].gonghao}}</text>
								</view>
								<view :style='{"padding":"0 0px","display":"inline-block"}'>
									<text class="icon iconfont icon-zan10" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#fff"}'></text>
									<text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"24rpx"}'>{{fangyuanxinxilist[5].thumbsupnum}}</text>
								</view>
								<view :style='{"padding":"0 20rpx","display":"inline-block"}'>
									<text class="icon iconfont icon-shoucang10" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#fff"}'></text>
									<text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"24rpx"}'>{{fangyuanxinxilist[5].storeupnum}}</text>
								</view>
								<view :style='{"padding":"0 0px","display":"inline-block"}'>
									<text class="icon iconfont icon-chakan2" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#fff"}'></text>
									<text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"24rpx"}'>{{fangyuanxinxilist[5].clicknum}}</text>
								</view>
							</view>
						</view>
					</view>
				</view>
			</view>
			<!-- 商品推荐 -->
			
			<!-- 商品列表 -->
			<!-- 商品列表 -->
			<!-- 新闻资讯 -->
			<view class="listBox news" :style='{"width":"100%","padding":"0 20rpx 40rpx 20rpx","margin":"40rpx 0 0","background":"linear-gradient( 135deg, #E4EFBE 0%, #EFE0BE 100%)","order":"3"}'>
				<view class="title" :style='{"padding":"0 0 0 60rpx","margin":"0","background":"url(http://codegen.caihongy.cn/20250210/97c3da7518f6421f973c9bfe967271e4.png) no-repeat center / 100% 100%","display":"flex","width":"100%","position":"relative","justifyContent":"flex-start","height":"80rpx"}'>
					<view :style='{"fontSize":"32rpx","lineHeight":"80rpx","color":"#FFFFFF","fontWeight":"500"}'>公告资讯</view>
					<view :style='{"alignItems":"center","top":"18rpx","color":"#6683A2","display":"flex","position":"absolute","right":"40rpx","justifyContent":"center"}' @tap="onPageTap('news')">
					  <text :style='{"color":"inherit","fontSize":"28rpx"}'>查看更多</text>
					  <text class="icon iconfont icon-gengduo1" :style='{"color":"inherit","fontSize":"28rpx"}'></text>
					</view>
				</view>
				<!-- 样式5 -->
				<view class="news-box2" :style='{"width":"100%","padding":"24rpx","margin":"0","background":"#fff","height":"auto"}'>
					<block v-for="(item,index) in news" :key="index">
						<view @tap="onNewsDetailTap(item.id)" v-if="index==0" class="list-item" :style='{"width":"100%","padding":"0","margin":"0 0 20rpx","position":"relative","height":"auto"}'>
							<image :style='{"width":"100%","objectFit":"cover","display":"block","height":"328rpx"}' mode="aspectFill" class="listmpic" :src="baseUrl+item.picture"></image>
							<view :style='{"width":"100%","background":"#B1BA85"}'>
								<view :style='{"padding":"0 20rpx","overflow":"hidden","whiteSpace":"nowrap","color":"#fff","lineHeight":"1.5","fontSize":"32rpx","textOverflow":"ellipsis"}' class="title">{{item.title}}</view>
								<view :style='{"padding":"0 20rpx","margin":"0","overflow":"hidden","color":"#fff","maxHeight":"80rpx","width":"100%","lineHeight":"40rpx","fontSize":"28rpx"}' class="text">{{item.introduction}}</view>
								<view :style='{"padding":"0 20rpx"}'>
									<text class="icon iconfont icon-shijian21" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#fff"}'></text>
									<text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"24rpx"}'>{{item.addtime.split(' ')[0].replace(/\-/g,'-')}}</text>
								</view>
								<view :style='{"padding":"0 20rpx","display":"inline-block"}'>
									<text class="icon iconfont icon-geren16" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#fff"}'></text>
									<text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"24rpx"}'>{{item.name}}</text>
								</view>
								<view :style='{"padding":"0 20rpx","display":"inline-block"}'>
									<text class="icon iconfont icon-zan10" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#fff"}'></text>
									<text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"24rpx"}'>{{item.thumbsupnum}}</text>
								</view>
								<view :style='{"padding":"0 20rpx","display":"inline-block"}'>
									<text class="icon iconfont icon-shoucang10" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#fff"}'></text>
									<text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"24rpx"}'>{{item.storeupnum}}</text>
								</view>
								<view :style='{"padding":"0 20rpx","display":"inline-block"}'>
									<text class="icon iconfont icon-chakan9" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#fff"}'></text>
									<text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"24rpx"}'>{{item.clicknum}}</text>
								</view>
							</view>
						</view>
						<view @tap="onNewsDetailTap(item.id)" v-if="index>0" class="list-item" :style='{"width":"100%","borderColor":"#ccc","margin":"0 0 20rpx 0","borderStyle":"solid","borderWidth":"0 0 2rpx 0","height":"auto"}'>
							<view :style='{"padding":"0 20rpx","overflow":"hidden","whiteSpace":"nowrap","color":"#333","lineHeight":"1.5","fontSize":"28rpx","textOverflow":"ellipsis"}' class="title">{{item.title}}</view>
							<view :style='{"padding":"0 20rpx","margin":"0","overflow":"hidden","color":"#666","maxHeight":"80rpx","width":"100%","lineHeight":"40rpx","fontSize":"28rpx"}' class="text">{{item.introduction}}</view>
							<view :style='{"padding":"0 20rpx"}'>
								<text class="icon iconfont icon-shijian21" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#666"}'></text>
								<text :style='{"color":"#666","lineHeight":"1.5","fontSize":"24rpx"}'>{{item.addtime.split(' ')[0].replace(/\-/g,'-')}}</text>
							</view>
							<view :style='{"padding":"0 20rpx","display":"inline-block"}'>
								<text class="icon iconfont icon-geren16" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#666"}'></text>
								<text :style='{"color":"#666","lineHeight":"1.5","fontSize":"24rpx"}'>{{item.name}}</text>
							</view>
							<view :style='{"padding":"0 20rpx","display":"inline-block"}'>
								<text class="icon iconfont icon-zan10" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#666"}'></text>
								<text :style='{"color":"#666","lineHeight":"1.5","fontSize":"24rpx"}'>{{item.thumbsupnum}}</text>
							</view>
							<view :style='{"padding":"0 20rpx","display":"inline-block"}'>
								<text class="icon iconfont icon-shoucang10" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#666"}'></text>
								<text :style='{"color":"#666","lineHeight":"1.5","fontSize":"24rpx"}'>{{item.storeupnum}}</text>
							</view>
							<view :style='{"padding":"0 20rpx","display":"inline-block"}'>
								<text class="icon iconfont icon-chakan9" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"24rpx","color":"#666"}'></text>
								<text :style='{"color":"#666","lineHeight":"1.5","fontSize":"24rpx"}'>{{item.clicknum}}</text>
							</view>
						</view>
					</block>
				</view>
			</view>
			<!-- 新闻资讯 -->
			<view v-if="scrollTop>200" @tap="scrollTopClick" :style='{"boxShadow":"0 8rpx 16rpx rgba(0,0,0,.3)","borderRadius":"50%","textAlign":"center","bottom":"20%","background":"#ff000030","width":"60rpx","lineHeight":"60rpx","position":"fixed","right":"20rpx","height":"60rpx"}'>
				<span class="icon iconfont icon-jiantou07" :style='{"color":"#fff"}'></span>
			</view>
		</view>
	</view>
</template>

<script>
    import menu from '@/utils/menu'
	import '@/assets/css/global-restaurant.css'
	import uniIcons from "@/components/uni-ui/lib/uni-icons/uni-icons.vue"
	export default {
		components: {
			uniIcons
		},
		data() {
			return {
				startX: 0,
				prevNumList4: '',
				numList4: 0,
				timerList4: null,
				flagList4: false,
				navigationActive: {"width":"16rpx","margin":"0 4rpx","borderRadius":"100%","background":"#F7EFDB","height":"16rpx"},
				navigationDefault: {"width":"16rpx","margin":"0 4rpx","borderRadius":"100%","background":"#fff","height":"16rpx"},
				options2: {
					effect: 'flip',
					loop : true
				},
				options3: {
					effect: 'cube',
					loop : true,
					cubeEffect: {
						shadow: true,
						slideShadows: true,
						shadowOffset: 20,
						shadowScale: 0.94,
					}
				},
				rows: 2,
				column: 4,
				iconArr: [
					'cuIcon-same',
					'cuIcon-deliver',
					'cuIcon-evaluate',
					'cuIcon-shop',
					'cuIcon-ticket',
					'cuIcon-cascades',
					'cuIcon-discover',
					'cuIcon-question',
					'cuIcon-pic',
					'cuIcon-filter',
					'cuIcon-footprint',
					'cuIcon-pulldown',
					'cuIcon-pullup',
					'cuIcon-moreandroid',
					'cuIcon-refund',
					'cuIcon-qrcode',
					'cuIcon-remind',
					'cuIcon-profile',
					'cuIcon-home',
					'cuIcon-message',
					'cuIcon-link',
					'cuIcon-lock',
					'cuIcon-unlock',
					'cuIcon-vip',
					'cuIcon-weibo',
					'cuIcon-activity',
					'cuIcon-friendadd',
					'cuIcon-friendfamous',
					'cuIcon-friend',
					'cuIcon-goods',
					'cuIcon-selection'
				],
				role : '',
				menuList: [],
				swiperMenuList:[],
				user: {},
				tableName:'',

				//轮播
				swiperList: [],
				fangyuanxinxilist: [],
				news: [],
				scrollTop: 0,
			}
		},
		onPageScroll(e) {
			this.scrollTop = e.scrollTop
		},
		watch: {
		},
		mounted() {
		},
		computed: {
			baseUrl() {
				return this.$base.url;
			},
		},
		async onLoad(){
		},
		async onShow() {
			if (this.timerList4&&this.timerList4!=null) clearInterval(this.timerList4)
			this.swiperMenuList = []
			this.role = uni.getStorageSync("appRole");
			let table = uni.getStorageSync("nowTable");
			let res = {}
			if(table) {
				res = await this.$api.session(table);
				this.user = res.data;
				this.tableName = table;
			}
			let menus = menu.list();
			this.menuList = menus;
			this.menuList.forEach((item,key) => {
				if(key==0) {
					item.frontMenu.forEach((item2,key2) => {
						if(item2.child[0].buttons.indexOf("查看")>-1) {
							this.swiperMenuList.push(item2);
						}
					})
				}
			})
			// let res;
			// 轮播图
			let swiperList = []
			res = await this.$api.list('config', {
				page: 1,
				limit: 5
			});
			for (let item of res.data.list) {
				if (item.name.indexOf('picture') >= 0 && item.value && item.value!="" && item.value!=null ) {
					swiperList.push({
						img: item.value,
						title: item.name,
						url: item.url
					});
				}
			}
			if (swiperList) {
				this.swiperList = swiperList;
			}
			
			this.prevNumList4 = this.swiperList.length - 1
			this.timerList4 = setInterval(this.autoPlayList4, 3000)

			// 推荐信息
			this.getRecommendList()
			this.getHomeList()
			this.getNewsList()
		},
		methods: {
			scrollTopClick(){
				uni.pageScrollTo({
					scrollTop: 0
				})
			},
			uGetRect(selector, all) {
				return new Promise(resolve => {
					uni.createSelectorQuery()
					.in(this)
					[all ? 'selectAll' : 'select'](selector)
					.boundingClientRect(rect => {
						if (all && Array.isArray(rect) && rect.length) {
							resolve(rect);
						}
						if (!all && rect) {
							resolve(rect);
						}
					})
					.exec();
				});
			},
			cloneData(data) {
				return JSON.parse(JSON.stringify(data));
			},
			newsTabClick2(index){
				this.newsIndex2 = index
				this.getNewsList()
			},
			async getNewsList(){
				let res;
				let params = {
					page: 1,
					limit: 4,
					sort: 'id',
					order: 'desc',
				}
				// 公告资讯
				res = await this.$api.list('news', params)
				this.news = res.data.list
			},
			homeTabClick2(index,name){
				this['home' + name + 'Index2'] = index
				this.getHomeList()
			},
			async getHomeList(){
				let res;
				let params;
			},
			recommendTabClick2(index,name){
				this[name + 'Index2'] = index
				this.getRecommendList()
			},
			async getRecommendList(){
				let res;
				let params;
				// 推荐信息
				params = {
					page: 1,
					limit: 6,
				}
				if(uni.getStorageSync("appUserid")) {
					res = await this.$api.recommend2('fangyuanxinxi', params);
				} else {
					res = await this.$api.recommend('fangyuanxinxi', params);
				}
				this.fangyuanxinxilist = res.data.list
				

			},
			autoPlayList4() {
				this.prevNumList4 = this.numList4
			
				this.numList4++
				if (this.numList4 == this.swiperList.length) this.numList4 = 0
			},
			touchStart(event) {
				this.startX = event.touches[0].clientX
				
				clearInterval(this.timerList4)
				this.flagList4 = true
			},
			touchMove(event) {
				const currentX = event.touches[0].clientX;
				const deltaX = currentX - this.startX;
				
				if (deltaX > 50) {
					// 向右滑动逻辑
					if (this.flagList4) {
						this.flagList4 = false
						
						this.prevNumList4 = this.numList4
						this.numList4++
						if (this.numList4 == this.swiperList.length) this.numList4 = 0
					}
					
				} else if (deltaX < -50) {
					// 向左滑动逻辑
					if (this.flagList4) {
						this.flagList4 = false
						
						this.prevNumList4 = this.numList4
						this.numList4--
						if (this.numList4 < 0) this.numList4 = this.swiperList.length - 1
					}
				}
			},
			touchEnd() {
				this.startX = 0
				this.timerList4 = setInterval(this.autoPlayList4, 3000)
				this.flagList4 = false
			},
			//轮播图跳转
			onSwiperTap(e) {
				if(e.url) {
					if (e.url.indexOf('https') != -1) {
						// #ifdef MP-WEIXIN
						uni.navigateTo({
						    url: '../../common/linkOthers/linkOthers?url=' + encodeURIComponent(e.url),
						});
						return false
						// #endif
						window.open(e.url)
					} else {
						this.$utils.jump(e.url)
					}
				}
			},
			// 新闻详情
			onNewsDetailTap(id) {
				this.$utils.jump(`../news-detail/news-detail?id=${id}`)
			},
			// 推荐列表点击详情
			onDetailTap(tableName, id) {
				this.$utils.jump(`../${tableName}/detail?id=${id}`)
			},
			onPageTap(tableName){
				uni.navigateTo({
					url: `../${tableName}/list`,
					fail: function(){
						uni.switchTab({
							url: `../${tableName}/list`
						});
					}
				});
				// this.$utils.jump(`../${tableName}/list`)
			},
			onPageTap2(index) {
				let url = '../' + index + '/list'
				if(index=='kefuguanli') {
					url = '../chat/chat'
				}
				uni.setStorageSync("useridTag",0);
				uni.navigateTo({
					url: url,
					fail: function() {
						uni.switchTab({
							url: url
						});
					}
				});
			}
		}
	}
</script>

<style lang="scss" scoped>
	.content {
		min-height: calc(100vh - 44px);
		box-sizing: border-box;
	}
	.list-swiper-4 .animate__animated {
		--animate-delay: 300ms;
	}
	.list {
		.style2 {
			.tabView {
				.tab {
					border: none;
					border-radius: 40rpx;
					padding: 0 30rpx;
					margin: 0 10rpx;
					color: #000;
					background: #EFE0BE;
				}
				.tabActive {
					border: none;
					border-radius: 40rpx;
					padding: 0 20rpx;
					margin: 0 10rpx;
					color: #fff;
					background: #B1BA85;
				}
			}
		}
	}

	.menu_view {
		padding: 20rpx;
		margin: 0;
		background: #fff;
		border-image: linear-gradient( 135deg, #E4EFBE 0%, #EFE0BE 100%) 1;
		display: flex;
		width: 100%;
		border-width: 20rpx;
		justify-content: flex-start;
		border-style: solid;
		flex-wrap: wrap;
		height: auto;
		.menu-item {
			padding: 12rpx 0;
			margin: 10rpx 0;
			color: #B1BA85;
			background: none;
			width: 25%;
			height: auto;
			.iconarr {
				border-radius: 100%;
				padding: 0;
				margin: 0px auto;
				color: inherit;
				background: none;
				display: block;
				width: 64rpx;
				font-size: 64rpx;
				line-height: 64rpx;
				height: 64rpx;
			}
			.text {
				padding: 0;
				margin: 12rpx auto 0;
				color: inherit;
				width: 100%;
				font-size: 28rpx;
				line-height: 28rpx;
				text-align: center;
			}
		}
	}
</style>
