<template>
<view class="content">
	<view :style='{"minHeight":"100%","width":"100%","padding":"0","position":"relative","background":"#FFFBF3","height":"auto"}'>
		<form :style='{"width":"100%","padding":"24rpx","background":"none","display":"block","height":"auto"}' class="app-update-pv">
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class="">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">房屋名称</view>
				<input :style='{"border":"none","padding":"0px 24rpx","margin":"0px","color":"rgb(0, 0, 0)","borderRadius":"40rpx","flex":"1","background":"rgba(255, 255, 255, 0)","fontSize":"28rpx","height":"80rpx"}' :disabled="ro.fangwumingcheng" v-model="ruleForm.fangwumingcheng" placeholder="房屋名称"  type="text"></input>
			</view>
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class="" @tap="fangwutupianTap">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">房屋图片</view>
				<image :style='{"width":"80rpx","borderRadius":"100%","objectFit":"cover","display":"block","height":"80rpx"}' class="avator" v-if="ruleForm.fangwutupian" :src="baseUrl+ruleForm.fangwutupian.split(',')[0]" mode="aspectFill"></image>
				<image :style='{"width":"80rpx","borderRadius":"100%","objectFit":"cover","display":"block","height":"80rpx"}' class="avator" v-else src="../../static/gen/upload.png" mode="aspectFill"></image>
			</view>
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class="">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">房屋户型</view>
				<input :style='{"border":"none","padding":"0px 24rpx","margin":"0px","color":"rgb(0, 0, 0)","borderRadius":"40rpx","flex":"1","background":"rgba(255, 255, 255, 0)","fontSize":"28rpx","height":"80rpx"}' :disabled="ro.fangwuhuxing" v-model="ruleForm.fangwuhuxing" placeholder="房屋户型"  type="text"></input>
			</view>
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class="">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">起始时间</view>
				<input :disabled="ro.qishishijian" :style='{"border":"none","padding":"0px 24rpx","margin":"0px","color":"rgb(0, 0, 0)","borderRadius":"40rpx","flex":"1","background":"rgba(255, 255, 255, 0)","fontSize":"28rpx","height":"80rpx"}' v-model="ruleForm.qishishijian" placeholder="起始时间" @tap="toggleTab('qishishijian')"></input>
			</view>
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class="">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">结束时间</view>
				<input :disabled="ro.jieshushijian" :style='{"border":"none","padding":"0px 24rpx","margin":"0px","color":"rgb(0, 0, 0)","borderRadius":"40rpx","flex":"1","background":"rgba(255, 255, 255, 0)","fontSize":"28rpx","height":"80rpx"}' v-model="ruleForm.jieshushijian" placeholder="结束时间" @tap="toggleTab('jieshushijian')"></input>
			</view>
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class="">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">租赁时间</view>
				<input :style='{"border":"none","padding":"0px 24rpx","margin":"0px","color":"rgb(0, 0, 0)","borderRadius":"40rpx","flex":"1","background":"rgba(255, 255, 255, 0)","fontSize":"28rpx","height":"80rpx"}' disabled v-model="zulinshijian" placeholder="租赁时间"></input>
			</view>
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class="">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">租金价格</view>
				<input :style='{"border":"none","padding":"0px 24rpx","margin":"0px","color":"rgb(0, 0, 0)","borderRadius":"40rpx","flex":"1","background":"rgba(255, 255, 255, 0)","fontSize":"28rpx","height":"80rpx"}' :disabled="ro.zujinjiage" v-model.number="ruleForm.zujinjiage" placeholder="租金价格" type="digit"></input>
			</view>
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class="">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">总计</view>
				<view :style='{"padding":"0px 24rpx","margin":"0px","lineHeight":"80rpx","fontSize":"28rpx","color":"rgb(0, 0, 0)","flex":"1"}' class="right-input">
					{{zongji}}
				</view>
			</view>
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class="">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">押金</view>
				<input :style='{"border":"none","padding":"0px 24rpx","margin":"0px","color":"rgb(0, 0, 0)","borderRadius":"40rpx","flex":"1","background":"rgba(255, 255, 255, 0)","fontSize":"28rpx","height":"80rpx"}' :disabled="ro.yajin" v-model.number="ruleForm.yajin" placeholder="押金" type="digit"></input>
			</view>
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class="">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">工号</view>
				<input :style='{"border":"none","padding":"0px 24rpx","margin":"0px","color":"rgb(0, 0, 0)","borderRadius":"40rpx","flex":"1","background":"rgba(255, 255, 255, 0)","fontSize":"28rpx","height":"80rpx"}' :disabled="ro.gonghao" v-model="ruleForm.gonghao" placeholder="工号"  type="text"></input>
			</view>
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class="">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">员工姓名</view>
				<input :style='{"border":"none","padding":"0px 24rpx","margin":"0px","color":"rgb(0, 0, 0)","borderRadius":"40rpx","flex":"1","background":"rgba(255, 255, 255, 0)","fontSize":"28rpx","height":"80rpx"}' :disabled="ro.yuangongxingming" v-model="ruleForm.yuangongxingming" placeholder="员工姓名"  type="text"></input>
			</view>
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class="">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">账号</view>
				<input :style='{"border":"none","padding":"0px 24rpx","margin":"0px","color":"rgb(0, 0, 0)","borderRadius":"40rpx","flex":"1","background":"rgba(255, 255, 255, 0)","fontSize":"28rpx","height":"80rpx"}' :disabled="ro.zhanghao" v-model="ruleForm.zhanghao" placeholder="账号"  type="text"></input>
			</view>
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class="">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">姓名</view>
				<input :style='{"border":"none","padding":"0px 24rpx","margin":"0px","color":"rgb(0, 0, 0)","borderRadius":"40rpx","flex":"1","background":"rgba(255, 255, 255, 0)","fontSize":"28rpx","height":"80rpx"}' :disabled="ro.xingming" v-model="ruleForm.xingming" placeholder="姓名"  type="text"></input>
			</view>
                 
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class="">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">租赁备注</view>
				<textarea :style='{"border":"none","padding":"24rpx","margin":"0px","color":"rgb(0, 0, 0)","borderRadius":"40rpx","flex":"1","background":"rgba(255, 255, 255, 0)","fontSize":"28rpx","height":"240rpx"}' v-model="ruleForm.zulinbeizhu" placeholder="租赁备注" :maxlength="-1"></textarea>
			</view>
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#ccc","borderWidth":"0 0 2rpx 0","width":"100%","borderStyle":"solid","height":"auto"}' class="">
				<view :style='{"width":"100%","lineHeight":"80rpx","fontSize":"28rpx","color":"#333","fontWeight":"500"}' class="title">租赁详情</view>
				<xia-editor ref="zulinxiangqing" :style='{"minHeight":"300rpx","border":"2rpx solid #efefef","width":"100%","height":"auto"}' v-model="ruleForm.zulinxiangqing" placeholder="租赁详情" @editorChange="zulinxiangqingChange"></xia-editor>
			</view>
			
			<view :style='{"width":"100%","flexWrap":"wrap","display":"flex","height":"auto"}' class="btn" >
				<button :style='{"border":"0","padding":"0px","margin":"0 0 20rpx 0","color":"#000","background":"linear-gradient( 135deg, #EFE0BE 0%, #E4EFBE 100%)","width":"100%","lineHeight":"80rpx","fontSize":"32rpx","fontWeight":"600","height":"80rpx"}' @tap="onSubmitTap" class="bg-red">提交</button>
			</view>
		</form>

		<w-picker  mode="dateTime" step="1" :current="false" :hasSecond="false" @confirm="qishishijianConfirm" ref="qishishijian" themeColor="#333333"></w-picker>
		<w-picker  mode="dateTime" step="1" :current="false" :hasSecond="false" @confirm="jieshushijianConfirm" ref="jieshushijian" themeColor="#333333"></w-picker>
	</view>
</view>
</template>

<script>
	import wPicker from "@/components/w-picker/w-picker.vue";
	import xiaEditor from '@/components/xia-editor/xia-editor';
	import multipleSelect from "@/components/momo-multipleSelect/momo-multipleSelect";
	export default {
		data() {
			return {
				cross:'',
				ruleForm: {
				fangwumingcheng: '',
				fangwutupian: '',
				fangwuhuxing: '',
				qishishijian: '',
				jieshushijian: '',
				zulinshijian: '',
				zujinjiage: '',
				zongji: '',
				yajin: '',
				jifen: '',
				zulinxiangqing: '',
				zulinbeizhu: '',
				gonghao: '',
				yuangongxingming: '',
				zhanghao: '',
				xingming: '',
				ispay: '',
				},
				// 登录用户信息
				user: {},
				ro:{
				   fangwumingcheng : false,
				   fangwutupian : false,
				   fangwuhuxing : false,
				   qishishijian : false,
				   jieshushijian : false,
				   zulinshijian : false,
				   zujinjiage : false,
				   zongji : false,
				   yajin : false,
				   jifen : false,
				   zulinxiangqing : false,
				   zulinbeizhu : false,
				   gonghao : false,
				   yuangongxingming : false,
				   zhanghao : false,
				   xingming : false,
				   ispay : false,
				},
				virtualPay: false,
			}
		},
		components: {
			wPicker,
			xiaEditor,
			multipleSelect,
		},
		computed: {
			zulinshijian : {
				get: function () {
					let d = this.ruleForm
					let a = 'd2.qishishijian-d2.jieshushijian'
					let n = a.split('-')
					let day = this.getFullDay(d[n[0].split('d2.')[1]], d[n[1].split('d2.')[1]])
					this.ruleForm.zulinshijian = day?day:0
					return day?day:0
				}
				
			},
			baseUrl() {
				return this.$base.url;
			},



			zongji:{
			get: function () {
				return 1*this.ruleForm.zulinshijian*this.ruleForm.zujinjiage+parseFloat(this.ruleForm.yajin)
			}
			},
		},
		async onLoad(options) {
			if(options.virtualPay){
				this.virtualPay = true
			}
			let table = uni.getStorageSync("nowTable");
			// 获取用户信息
			let res = await this.$api.session(table);
			this.user = res.data;
			
			// ss读取
			this.ruleForm.zhanghao = this.user.zhanghao
			this.ro.zhanghao = true;
			this.ruleForm.xingming = this.user.xingming
			this.ro.xingming = true;



			// 如果有登录，获取登录后保存的userid
			this.ruleForm.userid = uni.getStorageSync("appUserid")
			if (options.refid) {
				// 如果上一级页面传递了refid，获取改refid数据信息
				this.ruleForm.refid = Number(options.refid);
				this.ruleForm.nickname = uni.getStorageSync("nickname");
			}
			// 如果是更新操作
			if (options.id) {
				this.ruleForm.id = options.id;
				// 获取信息
				res = await this.$api.info(`fangwuzulin`, this.ruleForm.id);
				this.ruleForm = res.data;
			}
			// 跨表
			this.cross = options.cross;
			if(options.cross){
				var obj = uni.getStorageSync('crossObj');
				for (var o in obj){
					if(o=='fangwumingcheng'){
						this.ruleForm.fangwumingcheng = obj[o];
						this.ro.fangwumingcheng = true;
						continue;
					}
					if(o=='fangwutupian'){
						this.ruleForm.fangwutupian = obj[o].split(",")[0];
						this.ro.fangwutupian = true;
						continue;
					}
					if(o=='fangwuhuxing'){
						this.ruleForm.fangwuhuxing = obj[o];
						this.ro.fangwuhuxing = true;
						continue;
					}
					if(o=='qishishijian'){
						this.ruleForm.qishishijian = obj[o];
						this.ro.qishishijian = true;
						continue;
					}
					if(o=='jieshushijian'){
						this.ruleForm.jieshushijian = obj[o];
						this.ro.jieshushijian = true;
						continue;
					}
					if(o=='zulinshijian'){
						this.ruleForm.zulinshijian = obj[o];
						this.ro.zulinshijian = true;
						continue;
					}
					if(o=='zujinjiage'){
						this.ruleForm.zujinjiage = obj[o];
						this.ro.zujinjiage = true;
						continue;
					}
					if(o=='zongji'){
						this.ruleForm.zongji = obj[o];
						this.ro.zongji = true;
						continue;
					}
					if(o=='yajin'){
						this.ruleForm.yajin = obj[o];
						this.ro.yajin = true;
						continue;
					}
					if(o=='jifen'){
						this.ruleForm.jifen = obj[o];
						this.ro.jifen = true;
						continue;
					}
					if(o=='zulinxiangqing'){
						this.ruleForm.zulinxiangqing = obj[o];
						this.ro.zulinxiangqing = true;
						continue;
					}
					if(o=='zulinbeizhu'){
						this.ruleForm.zulinbeizhu = obj[o];
						this.ro.zulinbeizhu = true;
						continue;
					}
					if(o=='gonghao'){
						this.ruleForm.gonghao = obj[o];
						this.ro.gonghao = true;
						continue;
					}
					if(o=='yuangongxingming'){
						this.ruleForm.yuangongxingming = obj[o];
						this.ro.yuangongxingming = true;
						continue;
					}
					if(o=='zhanghao'){
						this.ruleForm.zhanghao = obj[o];
						this.ro.zhanghao = true;
						continue;
					}
					if(o=='xingming'){
						this.ruleForm.xingming = obj[o];
						this.ro.xingming = true;
						continue;
					}
				}
			}
			this.styleChange()
			this.$forceUpdate()
			if (uni.getStorageSync('raffleType') && uni.getStorageSync('raffleType') != null) {
				uni.removeStorageSync('raffleType')
				setTimeout(() => {
					this.onSubmitTap()
				}, 300)
			}
		},
		methods: {
			// 获取日期间隔 单位天
			getFullDay(first, last) {
				 // 验证输入是否为空
				    if (!first || !last) {
				        console.error('输入的日期不能为空');
				        return null;
				    }
				
				    // 将输入的日期字符串转换为 Date 对象
				    const date1 = new Date(first);
				    const date2 = new Date(last);
				
				    // 验证日期是否有效
				    if (isNaN(date1.getTime()) || isNaN(date2.getTime())) {
				        console.error('输入的日期格式无效，建议使用 YYYY/MM/DD 格式');
				        return null;
				    }
				
				    // 计算年份差和月份差
				    let yearsDiff = date2.getFullYear() - date1.getFullYear();
				    let monthsDiff = date2.getMonth() - date1.getMonth();
				
				    // 计算总月份差
				    let totalMonths = yearsDiff * 12 + monthsDiff;
				
				    // 如果结束日期的日小于开始日期的日，总月份数减 1
				    if (date2.getDate() < date1.getDate()) {
				        totalMonths--;
				    }
				
				    // 返回绝对值，确保结果为非负数
				    return Math.abs(totalMonths);
			},
			zulinxiangqingChange(e) {
				this.ruleForm.zulinxiangqing = e
			},
			styleChange() {
				this.$nextTick(()=>{
					// document.querySelectorAll('.app-update-pv . .uni-input-input').forEach(el=>{
					//   el.style.backgroundColor = this.addUpdateForm.input.content.backgroundColor
					// })
				})
			},

			// 多级联动参数


			// 日长控件选择日期时间
			qishishijianConfirm(val) {
				console.log(val)
				this.ruleForm.qishishijian = val.result;
				this.$forceUpdate();
			},
			// 日长控件选择日期时间
			jieshushijianConfirm(val) {
				console.log(val)
				this.ruleForm.jieshushijian = val.result;
				this.$forceUpdate();
			},


			fangwutupianTap() {
				let _this = this;
				if(this.ro.fangwutupian){
					return false
				}
				this.$api.upload(function(res) {
					_this.ruleForm.fangwutupian = 'upload/' + res.file;
					_this.$forceUpdate();
					_this.$nextTick(()=>{
						_this.styleChange()
					})
				});
			},

			getUUID () {
				return new Date().getTime();
			},
			async onSubmitTap() {
				let that = this
				if(!this.ruleForm.id) {
					this.ruleForm.ispay = '未支付'
				}
				if(this.ruleForm.zulinshijian==0){
					this.$utils.msg('租赁时间不能为空')
					return false
				}
				this.ruleForm.zongji = this.zongji
				//跨表计算判断
				var obj;
				if(this.ruleForm.zulinshijian&&(!this.$validate.isIntNumer(this.ruleForm.zulinshijian))){
					this.$utils.msg(`租赁时间应输入整数`);
					return
				}
				if(this.ruleForm.zujinjiage&&(!this.$validate.isNumber(this.ruleForm.zujinjiage))){
					this.$utils.msg(`租金价格应输入数字`);
					return
				}
				if(this.ruleForm.zongji&&(!this.$validate.isNumber(this.ruleForm.zongji))){
					this.$utils.msg(`总计应输入数字`);
					return
				}
				if(this.ruleForm.yajin&&(!this.$validate.isNumber(this.ruleForm.yajin))){
					this.$utils.msg(`押金应输入数字`);
					return
				}
				if(this.ruleForm.jifen&&(!this.$validate.isIntNumer(this.ruleForm.jifen))){
					this.$utils.msg(`积分应输入整数`);
					return
				}
				//更新跨表属性
				var crossuserid;
				var crossrefid;
				var crossoptnum;
				if(this.cross){
					var statusColumnName = uni.getStorageSync('statusColumnName');
					var statusColumnValue = uni.getStorageSync('statusColumnValue');
					if(statusColumnName!='') {
						if(!obj) {
							obj = uni.getStorageSync('crossObj');
						}
						if(!statusColumnName.startsWith("[")) {
							for (var o in obj){
								if(o==statusColumnName){
									obj[o] = statusColumnValue;
								}

							}
							var table = uni.getStorageSync('crossTable');
							await this.$api.update(`${table}`, obj);
						} else {
							   crossuserid=Number(uni.getStorageSync('appUserid'));
							   crossrefid=obj['id'];
							   crossoptnum=uni.getStorageSync('statusColumnName');
							   crossoptnum=crossoptnum.replace(/\[/,"").replace(/\]/,"");
						}
					}
				}
				if(crossrefid && crossuserid) {
					this.ruleForm.crossuserid=crossuserid;
					this.ruleForm.crossrefid=crossrefid;
					let params = {
						page: 1,
						limit:10,
						crossuserid:crossuserid,
						crossrefid:crossrefid,
					}
					let res = await this.$api.list(`fangwuzulin`, params);
					if (res.data.total >= crossoptnum) {
						this.$utils.msg(uni.getStorageSync('tips'));
						uni.removeStorageSync('crossCleanType');
						return false;
					} else {
				//跨表计算
						let oet = {}
						if(this.ruleForm.id){
							await this.$api.update(`fangwuzulin`, this.ruleForm);
						}else{
							oet = await this.$api.add(`fangwuzulin`, this.ruleForm);
						}
						if(this.isAuth('fangwuzulin','支付')&&this.cross) {
							uni.showModal({
								title: '提交成功！',
								content: '是否跳转支付？',
								success: async function(res) {
									if (res.confirm) {
										that.$utils.jump(`./detail?id=${oet.data}&userid=` + uni.getStorageSync("appUserid"))
									}else {
										if(that.cross){
											uni.setStorageSync('crossCleanType',true);
										}
										that.$utils.msgBack('提交成功');
									}
								}
							})
						}else {
							if(this.cross){
								uni.setStorageSync('crossCleanType',true);
							}
							this.$utils.msgBack('提交成功');
						}
					}
				} else {
				//跨表计算
					let oet = {}
					if(this.ruleForm.id){
						await this.$api.update(`fangwuzulin`, this.ruleForm);
					}else{
						oet = await this.$api.add(`fangwuzulin`, this.ruleForm);
					}
					if(this.isAuth('fangwuzulin','支付')&&this.cross) {
						uni.showModal({
							title: '提交成功！',
							content: '是否跳转支付？',
							success: async function(res) {
								if (res.confirm) {
									that.$utils.jump(`./detail?id=${oet.data}&userid=` + uni.getStorageSync("appUserid"))
								}else {
									if(that.cross){
										uni.setStorageSync('crossCleanType',true);
									}
									that.$utils.msgBack('提交成功');
								}
							}
						})
					}else {
						if(this.cross){
							uni.setStorageSync('crossCleanType',true);
						}
						this.$utils.msgBack('提交成功');
					}
				}
			},
			optionsChange(e) {
				this.index = e.target.value
			},
			bindDateChange(e) {
				this.date = e.target.value
			},
			getDate(type) {
				const date = new Date();
				let year = date.getFullYear();
				let month = date.getMonth() + 1;
				let day = date.getDate();
				if (type === 'start') {
					year = year - 60;
				} else if (type === 'end') {
					year = year + 2;
				}
				month = month > 9 ? month : '0' + month;;
				day = day > 9 ? day : '0' + day;
				return `${year}-${month}-${day}`;
			},
			toggleTab(str) {
				if(this.ro[str]){
					return false
				}
				this.$refs[str].show();
			},
		}
	}
</script>

<style lang="scss" scoped>
	.content {
		min-height: calc(100vh - 44px);
		box-sizing: border-box;
	}
</style>
