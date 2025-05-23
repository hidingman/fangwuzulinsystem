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
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class=" select">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">房屋户型</view>
				<picker :disabled="ro.fangwuhuxing" :style='{"width":"100%","flex":"1","height":"auto"}' @change="fangwuhuxingChange" :value="fangwuhuxingIndex" :range="fangwuhuxingOptions">
					<view :style='{"width":"100%","lineHeight":"80rpx","fontSize":"28rpx","color":"#000"}' class="uni-input">{{ruleForm.fangwuhuxing?ruleForm.fangwuhuxing:"请选择房屋户型"}}</view>
				</picker>
			</view>
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class="">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">房屋地址</view>
				<input :style='{"border":"none","padding":"0px 24rpx","margin":"0px","color":"rgb(0, 0, 0)","borderRadius":"40rpx","flex":"1","background":"rgba(255, 255, 255, 0)","fontSize":"28rpx","height":"80rpx"}' :disabled="ro.fangwudizhi" v-model="ruleForm.fangwudizhi" placeholder="房屋地址"  type="text"></input>
			</view>
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class="">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">面积</view>
				<input :style='{"border":"none","padding":"0px 24rpx","margin":"0px","color":"rgb(0, 0, 0)","borderRadius":"40rpx","flex":"1","background":"rgba(255, 255, 255, 0)","fontSize":"28rpx","height":"80rpx"}' :disabled="ro.mianji" v-model="ruleForm.mianji" placeholder="面积"  type="text"></input>
			</view>
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class="">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">租金价格</view>
				<input :style='{"border":"none","padding":"0px 24rpx","margin":"0px","color":"rgb(0, 0, 0)","borderRadius":"40rpx","flex":"1","background":"rgba(255, 255, 255, 0)","fontSize":"28rpx","height":"80rpx"}' :disabled="ro.zujinjiage" v-model.number="ruleForm.zujinjiage" placeholder="租金价格" type="digit"></input>
			</view>
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class=" select">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">押金方式</view>
				<picker :disabled="ro.yajinfangshi" :style='{"width":"100%","flex":"1","height":"auto"}' @change="yajinfangshiChange" :value="yajinfangshiIndex" :range="yajinfangshiOptions">
					<view :style='{"width":"100%","lineHeight":"80rpx","fontSize":"28rpx","color":"#000"}' class="uni-input">{{ruleForm.yajinfangshi?ruleForm.yajinfangshi:"请选择押金方式"}}</view>
				</picker>
			</view>
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class=" select">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">付款方式</view>
				<picker :disabled="ro.fukuanfangshi" :style='{"width":"100%","flex":"1","height":"auto"}' @change="fukuanfangshiChange" :value="fukuanfangshiIndex" :range="fukuanfangshiOptions">
					<view :style='{"width":"100%","lineHeight":"80rpx","fontSize":"28rpx","color":"#000"}' class="uni-input">{{ruleForm.fukuanfangshi?ruleForm.fukuanfangshi:"请选择付款方式"}}</view>
				</picker>
			</view>
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class="">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">房屋朝向</view>
				<input :style='{"border":"none","padding":"0px 24rpx","margin":"0px","color":"rgb(0, 0, 0)","borderRadius":"40rpx","flex":"1","background":"rgba(255, 255, 255, 0)","fontSize":"28rpx","height":"80rpx"}' :disabled="ro.fangwuchaoxiang" v-model="ruleForm.fangwuchaoxiang" placeholder="房屋朝向"  type="text"></input>
			</view>
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class="">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">押金</view>
				<input :style='{"border":"none","padding":"0px 24rpx","margin":"0px","color":"rgb(0, 0, 0)","borderRadius":"40rpx","flex":"1","background":"rgba(255, 255, 255, 0)","fontSize":"28rpx","height":"80rpx"}' :disabled="ro.yajin" v-model.number="ruleForm.yajin" placeholder="押金" type="digit"></input>
			</view>
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class="">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">小区</view>
				<input :style='{"border":"none","padding":"0px 24rpx","margin":"0px","color":"rgb(0, 0, 0)","borderRadius":"40rpx","flex":"1","background":"rgba(255, 255, 255, 0)","fontSize":"28rpx","height":"80rpx"}' :disabled="ro.xiaoqu" v-model="ruleForm.xiaoqu" placeholder="小区"  type="text"></input>
			</view>
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class="" @tap="fangwushipinTap">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">房屋视频</view>
				<input :style='{"border":"none","padding":"0px 24rpx","margin":"0px","color":"rgb(0, 0, 0)","borderRadius":"40rpx","flex":"1","background":"rgba(255, 255, 255, 0)","fontSize":"28rpx","height":"80rpx"}' v-if="ruleForm.fangwushipin"  v-model="baseUrl+ruleForm.fangwushipin" placeholder="房屋视频"></input>
				<image :style='{"width":"80rpx","borderRadius":"100%","objectFit":"cover","display":"block","height":"80rpx"}' class="avator" v-else src="../../static/gen/upload.png" mode="aspectFill"></image>
			</view>
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class="">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">楼栋单元</view>
				<input :style='{"border":"none","padding":"0px 24rpx","margin":"0px","color":"rgb(0, 0, 0)","borderRadius":"40rpx","flex":"1","background":"rgba(255, 255, 255, 0)","fontSize":"28rpx","height":"80rpx"}' :disabled="ro.loudongdanyuan" v-model="ruleForm.loudongdanyuan" placeholder="楼栋单元"  type="text"></input>
			</view>
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class="">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">房号</view>
				<input :style='{"border":"none","padding":"0px 24rpx","margin":"0px","color":"rgb(0, 0, 0)","borderRadius":"40rpx","flex":"1","background":"rgba(255, 255, 255, 0)","fontSize":"28rpx","height":"80rpx"}' :disabled="ro.fanghao" v-model="ruleForm.fanghao" placeholder="房号"  type="text"></input>
			</view>
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class="">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">房屋结构</view>
				<input :style='{"border":"none","padding":"0px 24rpx","margin":"0px","color":"rgb(0, 0, 0)","borderRadius":"40rpx","flex":"1","background":"rgba(255, 255, 255, 0)","fontSize":"28rpx","height":"80rpx"}' :disabled="ro.fangwujiegou" v-model="ruleForm.fangwujiegou" placeholder="房屋结构"  type="text"></input>
			</view>
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class=" select">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">房屋状态</view>
				<picker :disabled="ro.fangwuzhuangtai" :style='{"width":"100%","flex":"1","height":"auto"}' @change="fangwuzhuangtaiChange" :value="fangwuzhuangtaiIndex" :range="fangwuzhuangtaiOptions">
					<view :style='{"width":"100%","lineHeight":"80rpx","fontSize":"28rpx","color":"#000"}' class="uni-input">{{ruleForm.fangwuzhuangtai?ruleForm.fangwuzhuangtai:"请选择房屋状态"}}</view>
				</picker>
			</view>
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class="">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">房产证编号</view>
				<input :style='{"border":"none","padding":"0px 24rpx","margin":"0px","color":"rgb(0, 0, 0)","borderRadius":"40rpx","flex":"1","background":"rgba(255, 255, 255, 0)","fontSize":"28rpx","height":"80rpx"}' :disabled="ro.fangchanzhengbianhao" v-model="ruleForm.fangchanzhengbianhao" placeholder="房产证编号"  type="text"></input>
			</view>
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class="" @tap="fangchanzhengzhaopianTap">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">房产证照片</view>
				<image :style='{"width":"80rpx","borderRadius":"100%","objectFit":"cover","display":"block","height":"80rpx"}' class="avator" v-if="ruleForm.fangchanzhengzhaopian" :src="baseUrl+ruleForm.fangchanzhengzhaopian.split(',')[0]" mode="aspectFill"></image>
				<image :style='{"width":"80rpx","borderRadius":"100%","objectFit":"cover","display":"block","height":"80rpx"}' class="avator" v-else src="../../static/gen/upload.png" mode="aspectFill"></image>
			</view>
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class="">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">房主姓名</view>
				<input :style='{"border":"none","padding":"0px 24rpx","margin":"0px","color":"rgb(0, 0, 0)","borderRadius":"40rpx","flex":"1","background":"rgba(255, 255, 255, 0)","fontSize":"28rpx","height":"80rpx"}' :disabled="ro.fangzhuxingming" v-model="ruleForm.fangzhuxingming" placeholder="房主姓名"  type="text"></input>
			</view>
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class="">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">房主身份证</view>
				<input :style='{"border":"none","padding":"0px 24rpx","margin":"0px","color":"rgb(0, 0, 0)","borderRadius":"40rpx","flex":"1","background":"rgba(255, 255, 255, 0)","fontSize":"28rpx","height":"80rpx"}' :disabled="ro.fangzhushenfenzheng" v-model="ruleForm.fangzhushenfenzheng" placeholder="房主身份证"  type="text"></input>
			</view>
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class="">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">房主电话</view>
				<input :style='{"border":"none","padding":"0px 24rpx","margin":"0px","color":"rgb(0, 0, 0)","borderRadius":"40rpx","flex":"1","background":"rgba(255, 255, 255, 0)","fontSize":"28rpx","height":"80rpx"}' :disabled="ro.fangzhudianhua" v-model="ruleForm.fangzhudianhua" placeholder="房主电话"  type="text"></input>
			</view>
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class="">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">发布时间</view>
				<input :disabled="ro.fabushijian" :style='{"border":"none","padding":"0px 24rpx","margin":"0px","color":"rgb(0, 0, 0)","borderRadius":"40rpx","flex":"1","background":"rgba(255, 255, 255, 0)","fontSize":"28rpx","height":"80rpx"}' v-model="ruleForm.fabushijian" placeholder="发布时间" @tap="toggleTab('fabushijian')"></input>
			</view>
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class="">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">工号</view>
				<input :style='{"border":"none","padding":"0px 24rpx","margin":"0px","color":"rgb(0, 0, 0)","borderRadius":"40rpx","flex":"1","background":"rgba(255, 255, 255, 0)","fontSize":"28rpx","height":"80rpx"}' :disabled="ro.gonghao" v-model="ruleForm.gonghao" placeholder="工号"  type="text"></input>
			</view>
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#828282","alignItems":"center","borderWidth":"0 0 2rpx 0","display":"flex","width":"100%","borderStyle":"solid","height":"auto"}' class="">
				<view :style='{"width":"160rpx","padding":"0 20rpx 0 0","lineHeight":"80rpx","fontSize":"28rpx","color":"#7D7D7D","textAlign":"right"}' class="title">员工姓名</view>
				<input :style='{"border":"none","padding":"0px 24rpx","margin":"0px","color":"rgb(0, 0, 0)","borderRadius":"40rpx","flex":"1","background":"rgba(255, 255, 255, 0)","fontSize":"28rpx","height":"80rpx"}' :disabled="ro.yuangongxingming" v-model="ruleForm.yuangongxingming" placeholder="员工姓名"  type="text"></input>
			</view>
                                
			<view :style='{"padding":"12rpx 0","margin":"0 0 24rpx 0","borderColor":"#ccc","borderWidth":"0 0 2rpx 0","width":"100%","borderStyle":"solid","height":"auto"}' class="">
				<view :style='{"width":"100%","lineHeight":"80rpx","fontSize":"28rpx","color":"#333","fontWeight":"500"}' class="title">房屋详情</view>
				<xia-editor ref="fangwuxiangqing" :style='{"minHeight":"300rpx","border":"2rpx solid #efefef","width":"100%","height":"auto"}' v-model="ruleForm.fangwuxiangqing" placeholder="房屋详情" @editorChange="fangwuxiangqingChange"></xia-editor>
			</view>
			
			<view :style='{"width":"100%","flexWrap":"wrap","display":"flex","height":"auto"}' class="btn" >
				<button :style='{"border":"0","padding":"0px","margin":"0 0 20rpx 0","color":"#000","background":"linear-gradient( 135deg, #EFE0BE 0%, #E4EFBE 100%)","width":"100%","lineHeight":"80rpx","fontSize":"32rpx","fontWeight":"600","height":"80rpx"}' @tap="onSubmitTap" class="bg-red">提交</button>
			</view>
		</form>

		<w-picker  mode="dateTime" step="1" :current="false" :hasSecond="false" @confirm="fabushijianConfirm" ref="fabushijian" themeColor="#333333"></w-picker>
		<w-picker  mode="dateTime" step="1" :current="false" :hasSecond="false" @confirm="clicktimeConfirm" ref="clicktime" themeColor="#333333"></w-picker>
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
				fangwudizhi: '',
				mianji: '',
				zujinjiage: '',
				yajinfangshi: '',
				fukuanfangshi: '',
				fangwuchaoxiang: '',
				yajin: '',
				xiaoqu: '',
				fangwushipin: '',
				loudongdanyuan: '',
				fanghao: '',
					jifen: Number('5') ,
				fangwujiegou: '',
				fangwuzhuangtai: '',
				fangchanzhengbianhao: '',
				fangchanzhengzhaopian: '',
				fangzhuxingming: '',
				fangzhushenfenzheng: '',
				fangzhudianhua: '',
				fangwuxiangqing: '',
				fabushijian: '',
				gonghao: '',
				yuangongxingming: '',
				discussnum: '',
				storeupnum: '',
				},
				fangwuhuxingOptions: [],
				fangwuhuxingIndex: 0,
				yajinfangshiOptions: [],
				yajinfangshiIndex: 0,
				fukuanfangshiOptions: [],
				fukuanfangshiIndex: 0,
				fangwuzhuangtaiOptions: [],
				fangwuzhuangtaiIndex: 0,
				// 登录用户信息
				user: {},
				ro:{
				   fangwumingcheng : false,
				   fangwutupian : false,
				   fangwuhuxing : false,
				   fangwudizhi : false,
				   mianji : false,
				   zujinjiage : false,
				   yajinfangshi : false,
				   fukuanfangshi : false,
				   fangwuchaoxiang : false,
				   yajin : false,
				   xiaoqu : false,
				   fangwushipin : false,
				   loudongdanyuan : false,
				   fanghao : false,
				   jifen : false,
				   fangwujiegou : false,
				   fangwuzhuangtai : false,
				   fangchanzhengbianhao : false,
				   fangchanzhengzhaopian : false,
				   fangzhuxingming : false,
				   fangzhushenfenzheng : false,
				   fangzhudianhua : false,
				   fangwuxiangqing : false,
				   fabushijian : false,
				   gonghao : false,
				   yuangongxingming : false,
				   thumbsupnum : false,
				   crazilynum : false,
				   clicktime : false,
				   clicknum : false,
				   discussnum : false,
				   storeupnum : false,
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
			baseUrl() {
				return this.$base.url;
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
			this.ruleForm.gonghao = this.user.gonghao
			this.ro.gonghao = true;
			this.ruleForm.yuangongxingming = this.user.yuangongxingming
			this.ro.yuangongxingming = true;

			this.ro.jifen = true;

			// 下拉框
			res = await this.$api.option(`fangwuhuxing`,`fangwuhuxing`,{});
			this.fangwuhuxingOptions = res.data;
			this.fangwuhuxingOptions.unshift("请选择房屋户型");
			// 自定义下拉框值
			this.yajinfangshiOptions = "押一付一,押一付三,半年付,年付".split(',')
			// 自定义下拉框值
			this.fukuanfangshiOptions = "转账,现金".split(',')
			// 自定义下拉框值
			this.fangwuzhuangtaiOptions = "未租,已租".split(',')

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
				res = await this.$api.info(`fangyuanxinxi`, this.ruleForm.id);
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
					if(o=='fangwudizhi'){
						this.ruleForm.fangwudizhi = obj[o];
						this.ro.fangwudizhi = true;
						continue;
					}
					if(o=='mianji'){
						this.ruleForm.mianji = obj[o];
						this.ro.mianji = true;
						continue;
					}
					if(o=='zujinjiage'){
						this.ruleForm.zujinjiage = obj[o];
						this.ro.zujinjiage = true;
						continue;
					}
					if(o=='yajinfangshi'){
						this.ruleForm.yajinfangshi = obj[o];
						this.ro.yajinfangshi = true;
						continue;
					}
					if(o=='fukuanfangshi'){
						this.ruleForm.fukuanfangshi = obj[o];
						this.ro.fukuanfangshi = true;
						continue;
					}
					if(o=='fangwuchaoxiang'){
						this.ruleForm.fangwuchaoxiang = obj[o];
						this.ro.fangwuchaoxiang = true;
						continue;
					}
					if(o=='yajin'){
						this.ruleForm.yajin = obj[o];
						this.ro.yajin = true;
						continue;
					}
					if(o=='xiaoqu'){
						this.ruleForm.xiaoqu = obj[o];
						this.ro.xiaoqu = true;
						continue;
					}
					if(o=='fangwushipin'){
						this.ruleForm.fangwushipin = obj[o];
						this.ro.fangwushipin = true;
						continue;
					}
					if(o=='loudongdanyuan'){
						this.ruleForm.loudongdanyuan = obj[o];
						this.ro.loudongdanyuan = true;
						continue;
					}
					if(o=='fanghao'){
						this.ruleForm.fanghao = obj[o];
						this.ro.fanghao = true;
						continue;
					}
					if(o=='jifen'){
						this.ruleForm.jifen = obj[o];
						this.ro.jifen = true;
						continue;
					}
					if(o=='fangwujiegou'){
						this.ruleForm.fangwujiegou = obj[o];
						this.ro.fangwujiegou = true;
						continue;
					}
					if(o=='fangwuzhuangtai'){
						this.ruleForm.fangwuzhuangtai = obj[o];
						this.ro.fangwuzhuangtai = true;
						continue;
					}
					if(o=='fangchanzhengbianhao'){
						this.ruleForm.fangchanzhengbianhao = obj[o];
						this.ro.fangchanzhengbianhao = true;
						continue;
					}
					if(o=='fangchanzhengzhaopian'){
						this.ruleForm.fangchanzhengzhaopian = obj[o].split(",")[0];
						this.ro.fangchanzhengzhaopian = true;
						continue;
					}
					if(o=='fangzhuxingming'){
						this.ruleForm.fangzhuxingming = obj[o];
						this.ro.fangzhuxingming = true;
						continue;
					}
					if(o=='fangzhushenfenzheng'){
						this.ruleForm.fangzhushenfenzheng = obj[o];
						this.ro.fangzhushenfenzheng = true;
						continue;
					}
					if(o=='fangzhudianhua'){
						this.ruleForm.fangzhudianhua = obj[o];
						this.ro.fangzhudianhua = true;
						continue;
					}
					if(o=='fangwuxiangqing'){
						this.ruleForm.fangwuxiangqing = obj[o];
						this.ro.fangwuxiangqing = true;
						continue;
					}
					if(o=='fabushijian'){
						this.ruleForm.fabushijian = obj[o];
						this.ro.fabushijian = true;
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
					if(o=='thumbsupnum'){
						this.ruleForm.thumbsupnum = obj[o];
						this.ro.thumbsupnum = true;
						continue;
					}
					if(o=='crazilynum'){
						this.ruleForm.crazilynum = obj[o];
						this.ro.crazilynum = true;
						continue;
					}
					if(o=='clicktime'){
						this.ruleForm.clicktime = obj[o];
						this.ro.clicktime = true;
						continue;
					}
					if(o=='clicknum'){
						this.ruleForm.clicknum = obj[o];
						this.ro.clicknum = true;
						continue;
					}
					if(o=='discussnum'){
						this.ruleForm.discussnum = obj[o];
						this.ro.discussnum = true;
						continue;
					}
					if(o=='storeupnum'){
						this.ruleForm.storeupnum = obj[o];
						this.ro.storeupnum = true;
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
			fangwuxiangqingChange(e) {
				this.ruleForm.fangwuxiangqing = e
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
			fabushijianConfirm(val) {
				console.log(val)
				this.ruleForm.fabushijian = val.result;
				this.$forceUpdate();
			},
			// 日长控件选择日期时间
			clicktimeConfirm(val) {
				console.log(val)
				this.ruleForm.clicktime = val.result;
				this.$forceUpdate();
			},

			// 下拉变化
			fangwuhuxingChange(e) {
				this.fangwuhuxingIndex = e.target.value
				this.ruleForm.fangwuhuxing = this.fangwuhuxingOptions[this.fangwuhuxingIndex]
			},
			// 下拉变化
			yajinfangshiChange(e) {
				this.yajinfangshiIndex = e.target.value
				this.ruleForm.yajinfangshi = this.yajinfangshiOptions[this.yajinfangshiIndex]
			},
			// 下拉变化
			fukuanfangshiChange(e) {
				this.fukuanfangshiIndex = e.target.value
				this.ruleForm.fukuanfangshi = this.fukuanfangshiOptions[this.fukuanfangshiIndex]
			},
			// 下拉变化
			fangwuzhuangtaiChange(e) {
				this.fangwuzhuangtaiIndex = e.target.value
				this.ruleForm.fangwuzhuangtai = this.fangwuzhuangtaiOptions[this.fangwuzhuangtaiIndex]
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
			fangwushipinTap () {
				let _this = this;
				if(this.ro.fangwushipin){
					return false
				}
				this.$api.uploadMedia(function(res) {
					_this.ruleForm.fangwushipin = 'upload/' + res.file;
					_this.$forceUpdate();
				});
			},
			fangchanzhengzhaopianTap() {
				let _this = this;
				if(this.ro.fangchanzhengzhaopian){
					return false
				}
				this.$api.upload(function(res) {
					_this.ruleForm.fangchanzhengzhaopian = 'upload/' + res.file;
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
				//跨表计算判断
				var obj;
				if(this.ruleForm.zujinjiage&&(!this.$validate.isNumber(this.ruleForm.zujinjiage))){
					this.$utils.msg(`租金价格应输入数字`);
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
				if(this.ruleForm.fangzhushenfenzheng&&(!this.$validate.checkIdCard(this.ruleForm.fangzhushenfenzheng))){
					this.$utils.msg(`房主身份证应输入身份证格式`);
					return
				}
				if(this.ruleForm.fangzhudianhua&&(!this.$validate.isMobile(this.ruleForm.fangzhudianhua))){
					this.$utils.msg(`房主电话应输入手机格式`);
					return
				}
				if(this.ruleForm.thumbsupnum&&(!this.$validate.isIntNumer(this.ruleForm.thumbsupnum))){
					this.$utils.msg(`赞应输入整数`);
					return
				}
				if(this.ruleForm.crazilynum&&(!this.$validate.isIntNumer(this.ruleForm.crazilynum))){
					this.$utils.msg(`踩应输入整数`);
					return
				}
				if(this.ruleForm.clicknum&&(!this.$validate.isIntNumer(this.ruleForm.clicknum))){
					this.$utils.msg(`点击次数应输入整数`);
					return
				}
				if(this.ruleForm.discussnum&&(!this.$validate.isIntNumer(this.ruleForm.discussnum))){
					this.$utils.msg(`评论数应输入整数`);
					return
				}
				if(this.ruleForm.storeupnum&&(!this.$validate.isIntNumer(this.ruleForm.storeupnum))){
					this.$utils.msg(`收藏数应输入整数`);
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
					let res = await this.$api.list(`fangyuanxinxi`, params);
					if (res.data.total >= crossoptnum) {
						this.$utils.msg(uni.getStorageSync('tips'));
						uni.removeStorageSync('crossCleanType');
						return false;
					} else {
				//跨表计算
						let oet = {}
						if(this.ruleForm.id){
							await this.$api.update(`fangyuanxinxi`, this.ruleForm);
						}else{
							oet = await this.$api.add(`fangyuanxinxi`, this.ruleForm);
						}
						if(this.cross){
							uni.setStorageSync('crossCleanType',true);
						}
						this.$utils.msgBack('提交成功');
					}
				} else {
				//跨表计算
					let oet = {}
					if(this.ruleForm.id){
						await this.$api.update(`fangyuanxinxi`, this.ruleForm);
					}else{
						oet = await this.$api.add(`fangyuanxinxi`, this.ruleForm);
					}
					if(this.cross){
						uni.setStorageSync('crossCleanType',true);
					}
					this.$utils.msgBack('提交成功');
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
