<template>
	<view class="content">
		<view class="box" :style='{"width":"100%","padding":"0 60rpx","alignItems":"center","background":"linear-gradient( 135deg, #EFE0BE 0%, #E4EFBE 100%)","display":"flex","height":"100%"}'>
			<view :style='{"padding":"40rpx 20rpx","borderRadius":"20rpx","alignItems":"center","flexWrap":"wrap","background":"url(http://codegen.caihongy.cn/20250210/fdbb8d8faf4b4a62ad70343d4f4ff3fc.png) no-repeat center / 100% 100%","display":"flex","width":"100%","position":"relative","justifyContent":"center","height":"auto"}'>
				<image :style='{"width":"160rpx","margin":"0 auto 24rpx auto","borderRadius":"8rpx","display":"none","height":"160rpx"}' src="http://codegen.caihongy.cn/20201114/7856ba26477849ea828f481fa2773a95.jpg" mode="aspectFill"></image>
				<view v-if="loginType==1" :style='{"border":"none","width":"100%","margin":"0 0 40rpx 0","alignItems":"center","display":"flex","height":"auto"}' class="uni-form-item uni-column">
					<view :style='{"padding":"0","color":"#000000","borderRadius":"0","textAlign":"center","background":"none","flex":"none","width":"120rpx","lineHeight":"50rpx","fontSize":"28rpx","height":"50rpx"}' class="label">账号：</view>
					<input v-model="username" :style='{"padding":"0px 24rpx","margin":"0px","borderColor":"#DFDFDF","color":"#000","borderRadius":"0","flex":"1","background":"none","borderWidth":"0 0 2rpx 0","fontSize":"28rpx","borderStyle":"solid","height":"80rpx"}' type="text" class="uni-input" name="" placeholder="请输入账号" />
				</view>
				<view v-if="loginType==1" :style='{"border":"none","width":"100%","margin":"0 0 40rpx 0","alignItems":"center","display":"flex","height":"auto"}' class="uni-form-item uni-column">
					<view :style='{"padding":"0","color":"#000000","borderRadius":"0","textAlign":"center","background":"none","flex":"none","width":"120rpx","lineHeight":"50rpx","fontSize":"28rpx","height":"50rpx"}' class="label">密码：</view>
					<input v-model="password" password :style='{"padding":"0px 24rpx","margin":"0px","borderColor":"#DFDFDF","color":"#000","borderRadius":"0","flex":"1","background":"none","borderWidth":"0 0 2rpx 0","fontSize":"28rpx","borderStyle":"solid","height":"80rpx"}' type="password" class="uni-input" name="" placeholder="请输入密码" />
				</view>
				<view v-if="roleNum>1" :style='{"border":"none","width":"100%","margin":"0 0 40rpx 0","alignItems":"center","display":"flex","height":"auto"}'>
					<view :style='{"padding":"0","color":"#000000","borderRadius":"0","textAlign":"center","background":"none","flex":"none","width":"120rpx","lineHeight":"50rpx","fontSize":"28rpx","height":"50rpx"}' class="label">用户类型：</view>
					<picker @change="optionsChange" :value="index" :range="options" :style='{"padding":"0 20rpx","borderColor":"#DFDFDF","color":"#74868B","flex":"auto","borderWidth":"0 0 2rpx 0","width":"0","lineHeight":"80rpx","fontSize":"28rpx","borderStyle":"solid"}'>
						<view class="uni-picker-type">{{options[index]}}</view>
					</picker>
				</view>
				

				
				<button v-if="loginType==1" class="btn-submit" @tap="onLoginTap" type="primary" :style='{"border":"0","padding":"0px","margin":"280rpx 0 24rpx 0","color":"#000","borderRadius":"50%","background":"linear-gradient( 180deg, #EFE0BE 0%, #E4EFBE 100%)","width":"120rpx","lineHeight":"100rpx","fontSize":"32rpx","fontWeight":"600","height":"120rpx","order":"3"}'>登录</button>
				<button v-if="loginType==2" class="btn-submit" @tap="onFaceLoginTap" type="primary" :style='{"border":"0","padding":"0px","margin":"280rpx 0 24rpx 0","color":"#000","borderRadius":"50%","background":"linear-gradient( 180deg, #EFE0BE 0%, #E4EFBE 100%)","width":"120rpx","lineHeight":"100rpx","fontSize":"32rpx","fontWeight":"600","height":"120rpx","order":"3"}'>人脸识别登录</button>
				<view class="links" :style='{"width":"100%","margin":"0 0 24rpx 0","flexWrap":"wrap","display":"flex","height":"auto"}'>
					<view class="link-highlight" @tap="onRegisterTap('guke')" :style='{"color":"#74868B","padding":"0 8rpx","fontSize":"28rpx"}'>注册顾客</view>
				</view>
				
				<view class="idea1" :style='{"width":"100%","background":"red","display":"none","height":"80rpx"}'>idea1</view>
				<view class="idea2" :style='{"width":"100%","background":"red","display":"none","height":"80rpx"}'>idea2</view>
				<view class="idea3" :style='{"width":"100%","background":"red","display":"none","height":"80rpx"}'>idea3</view>
			</view>
		</view>
	</view>
</template>

<script>
	import menu from '@/utils/menu'
	export default {
		data() {
			return {
				username: '',
				password: '',
                loginType:1,
				codes: [{
				  num: 1,
				  color: '#000',
				  rotate: '10deg',
				  size: '16px'
				}, {
				  num: 2,
				  color: '#000',
				  rotate: '10deg',
				  size: '16px'
				}, {
				  num: 3,
				  color: '#000',
				  rotate: '10deg',
				  size: '16px'
				}, {
				  num: 4,
				  color: '#000',
				  rotate: '10deg',
				  size: '16px'
				}],
				options: [
					'请选择登录用户类型',
				],
                optionsValues: [
					'',
                    'guke',
				],
				index: 0,
				roleNum:0,

			}
		},
		onLoad() {
			let options = ['请选择登录用户类型'];
			let menus = menu.list();
			this.menuList = menus;
			for(let i=0;i<this.menuList.length;i++){
				if(this.menuList[i].hasFrontLogin=='是'){
					options.push(this.menuList[i].roleName);
					this.roleNum++;
				}
			}
			if(this.roleNum==1) {
				this.index = 1;
			}	
			this.options = options;
			this.styleChange()
		},
		onShow() {
		},
		mounted() {
		},
		methods: {
			styleChange() {
				this.$nextTick(()=>{
					// document.querySelectorAll('.uni-input .uni-input-input').forEach(el=>{
					//   el.style.backgroundColor = this.loginFrom.content.input.backgroundColor
					// })
				})
			},
			onRegisterTap(tableName) {
                uni.setStorageSync("loginTable", tableName);
				this.$utils.jump('../register/register')
			},
			async onLoginTap() {
                if (!this.username) {
					this.$utils.msg('请输入用户名')
					return
				}
                if (!this.password) {
					this.$utils.msg('请输入用户密码')
					return
				}
                if (!this.optionsValues[this.index]) {
					this.$utils.msg('请选择登录用户类型')
					return
				}

				this.loginPost()

			},
			async loginPost() {
				
				let res = await this.$api.login(`${this.optionsValues[this.index]}`, {
					username: this.username,
					password: this.password
				});
				uni.removeStorageSync("useridTag");
				uni.setStorageSync("appToken", res.token);
				uni.setStorageSync("nickname",this.username);
				uni.setStorageSync("nowTable", `${this.optionsValues[this.index]}`);
				res = await this.$api.session(`${this.optionsValues[this.index]}`);
				if(res.data.touxiang) {
				    uni.setStorageSync('headportrait', res.data.touxiang);
				} else if(res.data.headportrait) {
				    uni.setStorageSync('headportrait', res.data.headportrait);
				}
				uni.setStorageSync('userSession',JSON.stringify(res.data))
				// 保存用户id
				uni.setStorageSync("appUserid", res.data.id);
				if(res.data.vip) {
					uni.setStorageSync("vip", res.data.vip);
				}
				uni.setStorageSync("appRole", `${this.options[this.index]}`);
				this.$utils.tab('../index/index');
			},
			optionsChange(e) {
				this.index = e.target.value
			}
		}
	}
</script>

<style lang="scss" scoped>
	page {
		height: 100%;
	}
	
	.content {
		height: 100%;
		box-sizing: border-box;
	}
	
</style>
