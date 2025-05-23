<template>
	<view class="content">
		<view :style='{"minHeight":"100%","width":"100%","padding":"0","position":"relative","background":"#FFFBF3","height":"auto"}'>
			<view v-if="user&&user.id" :style='{"padding":"40rpx 24rpx","margin":"0 0 20rpx 0","background":"#E4EFBE","display":"flex","width":"100%","position":"relative","height":"auto"}' @tap="onPageTap('../user-info/user-info')" class="header" v-bind:class="{'status':isH5Plus}">
				<view :style='{"alignContent":"center","alignItems":"center","flexWrap":"wrap","flex":"1","display":"flex","justifyContent":"center","height":"100%"}' v-if="tableName=='guke'" class="userinfo">
					<image :style='{"width":"128rpx","padding":"0","margin":"0 20rpx 0 0","borderRadius":"100%","height":"128rpx"}' :src="user.touxiang?baseUrl+user.touxiang:require('../../static/gen/upload.png')"></image>
					<view :style='{"width":"auto","flexDirection":"column","justifyContent":"center","display":"flex"}' class="info">
						<view :style='{"width":"auto","lineHeight":"36rpx","fontSize":"28rpx","color":"#000"}'>{{user.zhanghao}}<text v-if="user.vip&& user.vip=='是'">(VIP)</text></view>
						<view :style='{"overflow":"hidden","whiteSpace":"nowrap","color":"#000","width":"auto","lineHeight":"36rpx","fontSize":"28rpx","textOverflow":"ellipsis"}'>{{user.xuqiu}}</view>
					</view>
					
					<view :style='{"margin":"20rpx 0 0","flexDirection":"row","display":"flex","gap":"20rpx","width":"100%","flexFlow":"wrap","justifyContent":"center"}' class="info">
						<view :style='{"padding":"8rpx","borderRadius":"12rpx","color":"#000","textAlign":"center","background":"#FFFFFF","flexDirection":"column","display":"flex","width":"30%","lineHeight":"36rpx","fontSize":"28rpx"}'>
							<span :style='{"color":"inherit"}'>手机：</span>
							<span :style='{"color":"inherit"}'>{{user.shouji}}</span>
						</view>
					</view>
					<view :style='{"width":"100%","margin":"40rpx 0 0","justifyContent":"center","display":"flex"}'>
					  <view :style='{"margin":"0","alignItems":"center","background":"#B1BA85","display":"flex","width":"50%","justifyContent":"center","height":"60rpx"}'>
					    <text class="icon iconfont icon-zan10" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"28rpx","color":"#fff"}'></text>
					    <text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"28rpx"}'>点赞 {{user.thumbsnum}}</text>
					  </view>
					  <view :style='{"margin":"0","alignItems":"center","background":"#EFE0BE","display":"flex","width":"50%","justifyContent":"center","height":"60rpx"}'>
					    <text class="icon iconfont icon-shoucang10" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"28rpx","color":"#8D8268"}'></text>
					    <text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"28rpx"}'>收藏 {{user.storenum}}</text>
					  </view>
					</view>
				</view>
				<view :style='{"alignContent":"center","alignItems":"center","flexWrap":"wrap","flex":"1","display":"flex","justifyContent":"center","height":"100%"}' v-if="tableName=='yuangong'" class="userinfo">
					<image :style='{"width":"128rpx","padding":"0","margin":"0 20rpx 0 0","borderRadius":"100%","height":"128rpx"}' :src="user.touxiang?baseUrl+user.touxiang:require('../../static/gen/upload.png')"></image>
					<view :style='{"width":"auto","flexDirection":"column","justifyContent":"center","display":"flex"}' class="info">
						<view :style='{"width":"auto","lineHeight":"36rpx","fontSize":"28rpx","color":"#000"}'>{{user.gonghao}}<text v-if="user.vip&& user.vip=='是'">(VIP)</text></view>
					</view>
					
					<view :style='{"margin":"20rpx 0 0","flexDirection":"row","display":"flex","gap":"20rpx","width":"100%","flexFlow":"wrap","justifyContent":"center"}' class="info">
						<view :style='{"padding":"8rpx","borderRadius":"12rpx","color":"#000","textAlign":"center","background":"#FFFFFF","flexDirection":"column","display":"flex","width":"30%","lineHeight":"36rpx","fontSize":"28rpx"}'>
							<span :style='{"color":"inherit"}'>电话：</span>
							<span :style='{"color":"inherit"}'>{{user.dianhua}}</span>
						</view>
					</view>
					<view :style='{"width":"100%","margin":"40rpx 0 0","justifyContent":"center","display":"flex"}'>
					  <view :style='{"margin":"0","alignItems":"center","background":"#B1BA85","display":"flex","width":"50%","justifyContent":"center","height":"60rpx"}'>
					    <text class="icon iconfont icon-zan10" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"28rpx","color":"#fff"}'></text>
					    <text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"28rpx"}'>点赞 {{user.thumbsnum}}</text>
					  </view>
					  <view :style='{"margin":"0","alignItems":"center","background":"#EFE0BE","display":"flex","width":"50%","justifyContent":"center","height":"60rpx"}'>
					    <text class="icon iconfont icon-shoucang10" :style='{"margin":"0 4rpx 0 0","lineHeight":"1.5","fontSize":"28rpx","color":"#8D8268"}'></text>
					    <text :style='{"color":"#fff","lineHeight":"1.5","fontSize":"28rpx"}'>收藏 {{user.storenum}}</text>
					  </view>
					</view>
				</view>
				<view :style='{"padding":"0 0 0 20rpx","alignItems":"center","top":"40rpx","flexDirection":"column","display":"flex","width":"84rpx","position":"absolute","right":"40rpx","justifyContent":"center","height":"auto"}' class="setting">
					<text class="icon iconfont icon-qita6" :style='{"border":"0","width":"48rpx","lineHeight":"48rpx","fontSize":"56rpx","color":"#000","borderRadius":"0"}'></text>
					<text :style='{"color":"#000","lineHeight":"2","fontSize":"28rpx"}'>设置</text>
				</view>
			</view>
			
			<view v-else :style='{"padding":"40rpx 24rpx","margin":"0 0 20rpx 0","background":"#E4EFBE","display":"flex","width":"100%","position":"relative","height":"auto"}' @tap="loginClick" class="header" v-bind:class="{'status':isH5Plus}">
				<view :style='{"alignItems":"center","color":"#fff","display":"flex","width":"100%","fontSize":"36rpx","justifyContent":"center","height":"100%"}'>登录/注册</view>
			</view>
			
			<view :style='{"width":"100%","background":"none","height":"auto"}' class="list">
				
				<view :style='{"width":"calc(100% - 40rpx)","boxShadow":"0px 4rpx 10rpx 0px rgba(0,0,0,0.3)","margin":"20rpx auto","borderRadius":"20rpx","background":"#fff","height":"auto"}'>
					<view :style='{"padding":"20rpx 0 0","margin":"0","color":"#fff","background":"url(http://codegen.caihongy.cn/20250208/5b4a70020ac84ac3b3aff15724223b68.png) no-repeat center / 100% 100%","display":"flex","width":"100%","letterSpacing":"2rpx","fontSize":"32rpx","position":"relative","justifyContent":"center","fontWeight":"600","height":"152rpx"}'>
						<view :style='{"fontSize":"32rpx","lineHeight":"72rpx","color":"#fff","fontWeight":"600"}'>我的服务</view>
					</view>
					<view class="me-menu-view">
						<block v-for="item in menuList" v-bind:key="item.roleName">
							<block v-if="role==item.roleName" v-bind:key="menu.menu" v-for="(menu,index) in item.backMenu">
								<block v-bind:key="child.tableName" v-for=" (child,sort) in menu.child">
									<view class="me-menu-item" v-if="child.tableName!='orders/已发货' && child.tableName!='orders/已退款' &&child.tableName!='orders/已取消' && child.tableName!='orders/未支付' && child.tableName!='orders/已支付' && child.tableName!='orders/已完成' && child.tableName!='exampaper' && child.tableName!='examquestion'&& child.tableName!='hasTranslate' " @tap="onPageTap('../'+child.tableName+'/list?userid='+user.id+'&menuJump='+child.menuJump)" hover-class="hover">
										<view class="me-menu-icon" :class="child.appFrontIcon"></view>
										<view class="text">{{child.menu}}</view>
										<view class="icon iconfont icon-jinru"></view>
									</view>
								</block>
							</block>
						</block>


						<view v-if="user&&user.id" @tap="onPageTap('../chat/chat')" class="me-menu-item" hover-class="hover">
							<view class="cuIcon-service me-menu-icon"></view>
							<view class="text">智能客服</view>
							<view class="icon iconfont icon-jinru"></view>
						</view>
						<view v-if="user&&user.id" @tap="passwordShow()" class="me-menu-item" hover-class="hover">
							<view class="cuIcon-lock me-menu-icon"></view>
							<view class="text">修改密码</view>
							<view class="icon iconfont icon-jinru"></view>
						</view>
					</view>
				</view>
			</view>
			
			<!-- 推荐 -->
			<view class="recommend" :style='{"width":"100%","padding":"0 0 20rpx 0","margin":"40rpx 0 0","background":"linear-gradient( 180deg, #E4EFBE 0%, #EFE0BE 100%)","height":"auto"}' v-if="recommendList.length">
				<view :style='{"padding":"8rpx 0 0","margin":"0","color":"#fff","background":"url(http://codegen.caihongy.cn/20250210/f6078ca119e2499f82ca68a92f6a1baa.png) no-repeat center / 100% 100%","display":"flex","width":"100%","fontSize":"32rpx","position":"relative","justifyContent":"center","height":"80rpx"}'>为你推荐</view>
				<view :style='{"padding":"20rpx","margin":"20rpx auto 0","borderRadius":"20rpx","flexWrap":"wrap","background":"#fff","display":"flex","width":"calc(100% - 40rpx)","justifyContent":"space-between","height":"auto"}'>
					<view v-for="(item,index) in recommendList" :key="index" :style='{"border":"2rpx solid #EFE0BE","padding":"0","alignItems":"center","borderRadius":"20rpx","background":"#fff","flexDirection":"column","display":"flex","width":"48%","position":"relative","justifyContent":"center","height":"auto"}' @click="recommendDetail(item.id)">
						<image v-if="preHttp(item[recommendPicture])" :style='{"width":"100%","objectFit":"cover","borderRadius":"20rpx","display":"block","height":"160rpx"}' :src="item[recommendPicture]"></image>
						<image v-else :style='{"width":"100%","objectFit":"cover","borderRadius":"20rpx","display":"block","height":"160rpx"}' :src="item[recommendPicture]?(baseUrl + item[recommendPicture].split(',')[0]):''"></image>
						<view :style='{"padding":"0","overflow":"hidden","whiteSpace":"nowrap","color":"#010F33","maxHeight":"80rpx","textAlign":"center","width":"100%","lineHeight":"40rpx","fontSize":"28rpx","textOverflow":"ellipsis"}'>{{item[recommendTitle]}}</view>
						<view :style='{"padding":"12rpx","lineHeight":"24rpx","fontSize":"24rpx","color":"#959595","textAlign":"center","background":"none"}'>{{item.addtime}}</view>
					</view>
				</view>
			</view>
			<view style="width: 100%;height: 80px"></view>
		</view>
		<uni-popup class="popup-content" ref="passwordPopup" type="bottom">
			<view class="passwordForm">
				<view class="passwordInput">
					<input type="password" v-model="passwordForm.mima" placeholder="原密码" />
				</view>
				<view class="passwordInput">
					<input type="password" v-model="passwordForm.newmima" placeholder="新密码" />
				</view>
				<view class="passwordInput">
					<input type="password" v-model="passwordForm.newmima1" placeholder="确认密码" />
				</view>
				<view class="passwordBtnView">
					<button class="passwordBtn" @click="updatePassword">更新</button>
				</view>
			</view>
		</uni-popup>
	</view>
</template>
<script>
	import menu from '@/utils/menu'
	export default {
		data() {
			return {
				isH5Plus: true,
				user: {},
				tableName:'',
				role: '',
				menuList: [],
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
				recommendList: [],
				recommendTable: '',
				recommendTitle: '',
				recommendPicture: '',
				passwordForm: {
					mima: '',
					newmima: '',
					newmima1: '',
				},
			};
		},
		computed: {
			baseUrl() {
				return this.$base.url;
			},
		},
		async onLoad(){
			let menus = menu.list();
			this.menuList = menus;
		},
		async onShow(){
            uni.removeStorageSync("useridTag");
			this.role = uni.getStorageSync("appRole");
			await this.getSession()
			this.getThumbsup()
			this.getStoreup()
			this.tableName = uni.getStorageSync("nowTable");
			let menus = menu.list();
			this.menuList = menus;
			this.getRecommendList()
			this.$forceUpdate()
		},
		methods: {
			async getSession(){
				let table = uni.getStorageSync("nowTable");
				let res = await this.$api.session(table);
				this.user = res.data;
				this.$forceUpdate()
			},
			async getThumbsup(){
				if(!this.user||!this.user.id){
					return false
				}
				let res = await this.$api.page('storeup',{userid: this.user.id,type: 21})
				this.user.thumbsnum = Number(res.data.total)
				this.$forceUpdate()
			},
			async getStoreup(){
				if(!this.user||!this.user.id){
					return false
				}
				let res = await this.$api.page('storeup',{userid: this.user.id,type: 1})
				this.user.storenum = Number(res.data.total)
				this.$forceUpdate()
			},
			loginClick(){
				uni.navigateTo({
					url: '../login/login'
				});
			},
			onPageTap(url) {
                uni.setStorageSync("useridTag",1);
				uni.navigateTo({
					url: url,
					fail: function() {
						uni.switchTab({
							url: url
						});
					}
				});
			},
			hasTranslate(){
				for(let x in this.menuList){
					if(this.menuList[x].roleName == this.role){
						for(let i in this.menuList[x].backMenu){
							if(this.menuList[x].backMenu[i].child.length) {
								for(let j in this.menuList[x].backMenu[i].child) {
									if(this.menuList[x].backMenu[i].child[j].tableName=='hasTranslate'){
										return true
									}
								}
							}
						}
					}
				}
				return false
			},
			
			preHttp(str) {
			    return str && str.substr(0,4)=='http';
			},
			async getRecommendList(){
				let params = {
					page: 1,
					limit: 4,
				}
				let res;
				if(uni.getStorageSync("appUserid")) {
					res = await this.$api.recommend2('fangyuanxinxi', params);
				} else {
					res = await this.$api.recommend('fangyuanxinxi', params);
				}
				this.recommendList = res.data.list
				this.recommendTable = 'fangyuanxinxi'
				this.recommendTitle = 'fangwumingcheng'
				this.recommendPicture = 'fangwutupian'
				this.$forceUpdate()
			},
			recommendDetail(id) {
				uni.navigateTo({
					url: `../${this.recommendTable}/detail?id=${id}`,
					fail: function() {
						uni.switchTab({
							url: `../${this.recommendTable}/detail?id=${id}`
						});
					}
				});
			},
			passwordShow() {
				this.passwordForm = {
					mima: '',
					newmima: '',
					newmima1: '',
				}
				this.$forceUpdate()
				this.$refs.passwordPopup.open()
			},
			async updatePassword() {
				if (this.passwordForm.mima == ''){
					this.$utils.msg('原密码不能为空')
					return false
				}
				if (this.passwordForm.newmima == ''){
					this.$utils.msg('新密码不能为空')
					return false
				}
				if (this.passwordForm.newmima1 == ''){
					this.$utils.msg('确认密码不能为空')
					return false
				}
				let password = ''
				if (this.user.mima) {
					password = this.user.mima;
				} else if (this.user.password) {
					password = this.user.password;
				}
				if (this.tableName == 'guke') {
				}
				let newpassword = this.passwordForm.mima
				if(password != newpassword){
					this.$utils.msg('原密码不正确')
					return false
				}
				if (this.passwordForm.newmima != this.passwordForm.newmima1){
					this.$utils.msg('两次密码不一致')
					return false
				}
				if (this.passwordForm.mima == this.passwordForm.newmima){
					this.$utils.msg('新密码与原密码相同！')
					return false
				}
				this.user.password = this.passwordForm.newmima
				this.user.mima = this.passwordForm.newmima
				await this.$api.update(this.tableName,this.user)
				this.$utils.msg('修改密码成功,下次登录系统生效')
				this.$refs.passwordPopup.close()
				this.getSession()
			},
		}
	}
</script>

<style lang="scss" scoped>
	.content {
		height: calc(100vh - 94px);
		box-sizing: border-box;
	}
	.passwordForm {
		width: 100%;
		padding: 10px 20px 60px;
		background: #fff;
	
		.passwordInput {
			background: #eee;
			margin: 0 0 10px;
			padding: 0 10px;
	
			input {
				height: 36px;
				line-height: 36px;
				font-size: 14px;
			}
		}
	
		.passwordBtnView {
			width: 100%;
			display: flex;
			align-items: center;
			justify-content: center;
			padding: 10px 0;
	
			.passwordBtn {
				padding: 0 20px;
				width: auto;
				height: 20px;
				line-height: 20px;
				font-size: 14px;
				color: #fff;
				background: #007aff;
				border: none;
				border-radius: 5px;
	
			}
		}
	}
	.me-menu-view {
		padding: 0 24rpx;
		background: #fff;
		display: flex;
		width: 100%;
		flex-wrap: wrap;
		height: auto;
		.me-menu-item {
			display: flex;
			width: 100%;
			border-color: #ccc;
			border-width: 0 0 2rpx;
			align-items: center;
			border-style: solid;
			height: auto;
			.me-menu-icon {
				color: rgba(0, 186, 189, 1);
				font-size: 40rpx;
				line-height: 1;
			}
			.text {
				padding: 0 20rpx;
				color: #333;
				flex: 1;
				width: 100%;
				font-size: 28rpx;
				line-height: 88rpx;
			}
			.icon {
				color: #999;
				width: 28rpx;
				font-size: 28rpx;
				line-height: 28rpx;
			}
		}
	}
</style>
