<template>
<view class="content">
	<view class="box" :style='{"minHeight":"100%","padding":"0 60rpx","alignItems":"flex-start","background":"linear-gradient( 135deg, #EFE0BE 0%, #E4EFBE 100%)","display":"flex","width":"100%","height":"auto"}'>
		<view :style='{"padding":"40rpx 20rpx","margin":"40rpx 0","borderRadius":"20rpx","background":"#fff","width":"100%","position":"relative","height":"auto"}'>
			<image :style='{"width":"160rpx","margin":"0 auto 24rpx auto","borderRadius":"8rpx","display":"none","height":"160rpx"}' src="http://codegen.caihongy.cn/20201114/7856ba26477849ea828f481fa2773a95.jpg" mode="aspectFill"></image>
			<view :style='{"border":"none","width":"100%","margin":"0 0 40rpx 0","alignItems":"center","display":"flex","height":"auto"}' v-if="tableName=='guke'" class="uni-form-item uni-column">
				<view :style='{"padding":"0","color":"#000000","borderRadius":"0","textAlign":"center","background":"none","flex":"none","width":"auto","lineHeight":"50rpx","fontSize":"28rpx","minWidth":"160rpx","height":"50rpx"}' class="label">账号：</view>
				<input :style='{"padding":"0px 24rpx","margin":"0px","borderColor":"#DFDFDF","color":"#000","borderRadius":"0","flex":"1","background":"none","borderWidth":"0 0 2rpx 0","fontSize":"28rpx","borderStyle":"solid","height":"80rpx"}'  v-model="ruleForm.zhanghao"  type="text"  class="uni-input" name="" placeholder="账号" />
			</view>
			<view :style='{"border":"none","width":"100%","margin":"0 0 40rpx 0","alignItems":"center","display":"flex","height":"auto"}' v-if="tableName=='guke'" class="uni-form-item uni-column">
				<view :style='{"padding":"0","color":"#000000","borderRadius":"0","textAlign":"center","background":"none","flex":"none","width":"auto","lineHeight":"50rpx","fontSize":"28rpx","minWidth":"160rpx","height":"50rpx"}' class="label">密码：</view>
				<input :style='{"padding":"0px 24rpx","margin":"0px","borderColor":"#DFDFDF","color":"#000","borderRadius":"0","flex":"1","background":"none","borderWidth":"0 0 2rpx 0","fontSize":"28rpx","borderStyle":"solid","height":"80rpx"}'  v-model="ruleForm.mima" type="password"  class="uni-input" name="" placeholder="密码" />
			</view>
			<view :style='{"border":"none","width":"100%","margin":"0 0 40rpx 0","alignItems":"center","display":"flex","height":"auto"}' v-if="tableName=='guke'" class="uni-form-item uni-column">
				<view :style='{"padding":"0","color":"#000000","borderRadius":"0","textAlign":"center","background":"none","flex":"none","width":"auto","lineHeight":"50rpx","fontSize":"28rpx","minWidth":"160rpx","height":"50rpx"}' class="label">确认密码：</view>
				<input :style='{"padding":"0px 24rpx","margin":"0px","borderColor":"#DFDFDF","color":"#000","borderRadius":"0","flex":"1","background":"none","borderWidth":"0 0 2rpx 0","fontSize":"28rpx","borderStyle":"solid","height":"80rpx"}' v-model="ruleForm.mima2" type="password" class="uni-input" name="" placeholder="确认密码" />
			</view>
			<view :style='{"border":"none","width":"100%","margin":"0 0 40rpx 0","alignItems":"center","display":"flex","height":"auto"}' v-if="tableName=='guke'" class="uni-form-item uni-column">
				<view :style='{"padding":"0","color":"#000000","borderRadius":"0","textAlign":"center","background":"none","flex":"none","width":"auto","lineHeight":"50rpx","fontSize":"28rpx","minWidth":"160rpx","height":"50rpx"}' class="label">姓名：</view>
				<input :style='{"padding":"0px 24rpx","margin":"0px","borderColor":"#DFDFDF","color":"#000","borderRadius":"0","flex":"1","background":"none","borderWidth":"0 0 2rpx 0","fontSize":"28rpx","borderStyle":"solid","height":"80rpx"}'  v-model="ruleForm.xingming"  type="text"  class="uni-input" name="" placeholder="姓名" />
			</view>
            <view :style='{"border":"none","width":"100%","margin":"0 0 40rpx 0","alignItems":"center","display":"flex","height":"auto"}' v-if="tableName=='guke'" class="">
				<view :style='{"padding":"0","color":"#000000","borderRadius":"0","textAlign":"center","background":"none","flex":"none","width":"auto","lineHeight":"50rpx","fontSize":"28rpx","minWidth":"160rpx","height":"50rpx"}' class="label">性别：</view>
				<picker :style='{"width":"100%","margin":"0","display":"flex","height":"auto"}' v-if="tableName=='guke'"  @change="gukexingbieChange" :value="gukexingbieIndex" :range="gukexingbieOptions">
					<view>
						<view :style='{"padding":"0 20rpx","lineHeight":"80rpx","fontSize":"28rpx","color":"#74868B","flex":"1"}' class="uni-input">{{ruleForm.xingbie?ruleForm.xingbie:"请选择性别"}}</view>
					</view>
				</picker>
			</view>
			<view :style='{"border":"none","width":"100%","margin":"0 0 40rpx 0","alignItems":"center","display":"flex","height":"auto"}' v-if="tableName=='guke'" class="uni-form-item uni-column">
				<view :style='{"padding":"0","color":"#000000","borderRadius":"0","textAlign":"center","background":"none","flex":"none","width":"auto","lineHeight":"50rpx","fontSize":"28rpx","minWidth":"160rpx","height":"50rpx"}' class="label">手机：</view>
				<input :style='{"padding":"0px 24rpx","margin":"0px","borderColor":"#DFDFDF","color":"#000","borderRadius":"0","flex":"1","background":"none","borderWidth":"0 0 2rpx 0","fontSize":"28rpx","borderStyle":"solid","height":"80rpx"}'  v-model="ruleForm.shouji"  type="text"  class="uni-input" name="" placeholder="手机" />
			</view>
            <view :style='{"border":"none","width":"100%","margin":"0 0 40rpx 0","alignItems":"center","display":"flex","height":"auto"}' v-if="tableName=='guke'" @tap="guketouxiangTap" class="">
				<view :style='{"padding":"0","color":"#000000","borderRadius":"0","textAlign":"center","background":"none","flex":"none","width":"auto","lineHeight":"50rpx","fontSize":"28rpx","minWidth":"160rpx","height":"50rpx"}' class="label">头像：</view>
				<image :style='{"width":"120rpx","borderRadius":"16rpx","display":"block","height":"120rpx"}' v-if="ruleForm.touxiang" class="avator" :src="baseUrl+ruleForm.touxiang" mode=""></image>
                <image :style='{"width":"120rpx","borderRadius":"16rpx","display":"block","height":"120rpx"}' v-else class="avator" src="../../static/gen/upload.png" mode=""></image>
            </view>
			<view :style='{"border":"none","width":"100%","margin":"0 0 40rpx 0","alignItems":"center","display":"flex","height":"auto"}' v-if="tableName=='guke'" class="uni-form-item uni-column">
				<view :style='{"padding":"0","color":"#000000","borderRadius":"0","textAlign":"center","background":"none","flex":"none","width":"auto","lineHeight":"50rpx","fontSize":"28rpx","minWidth":"160rpx","height":"50rpx"}' class="label">身份证：</view>
				<input :style='{"padding":"0px 24rpx","margin":"0px","borderColor":"#DFDFDF","color":"#000","borderRadius":"0","flex":"1","background":"none","borderWidth":"0 0 2rpx 0","fontSize":"28rpx","borderStyle":"solid","height":"80rpx"}'  v-model="ruleForm.shenfenzheng"  type="text"  class="uni-input" name="" placeholder="身份证" />
			</view>
			<view :style='{"border":"none","width":"100%","margin":"0 0 40rpx 0","alignItems":"center","display":"flex","height":"auto"}' v-if="tableName=='guke'" class="uni-form-item uni-column">
				<view :style='{"padding":"0","color":"#000000","borderRadius":"0","textAlign":"center","background":"none","flex":"none","width":"auto","lineHeight":"50rpx","fontSize":"28rpx","minWidth":"160rpx","height":"50rpx"}' class="label">民族：</view>
				<input :style='{"padding":"0px 24rpx","margin":"0px","borderColor":"#DFDFDF","color":"#000","borderRadius":"0","flex":"1","background":"none","borderWidth":"0 0 2rpx 0","fontSize":"28rpx","borderStyle":"solid","height":"80rpx"}'  v-model="ruleForm.minzu"  type="text"  class="uni-input" name="" placeholder="民族" />
			</view>
			<view :style='{"border":"none","width":"100%","margin":"0 0 40rpx 0","alignItems":"center","display":"flex","height":"auto"}' v-if="tableName=='guke'" class="uni-form-item uni-column">
				<view :style='{"padding":"0","color":"#000000","borderRadius":"0","textAlign":"center","background":"none","flex":"none","width":"auto","lineHeight":"50rpx","fontSize":"28rpx","minWidth":"160rpx","height":"50rpx"}' class="label">籍贯：</view>
				<input :style='{"padding":"0px 24rpx","margin":"0px","borderColor":"#DFDFDF","color":"#000","borderRadius":"0","flex":"1","background":"none","borderWidth":"0 0 2rpx 0","fontSize":"28rpx","borderStyle":"solid","height":"80rpx"}'  v-model="ruleForm.jiguan"  type="text"  class="uni-input" name="" placeholder="籍贯" />
			</view>
            <view :style='{"border":"none","width":"100%","margin":"0 0 40rpx 0","alignItems":"center","display":"flex","height":"auto"}' v-if="tableName=='guke'" class="">
				<view :style='{"padding":"0","color":"#000000","borderRadius":"0","textAlign":"center","background":"none","flex":"none","width":"auto","lineHeight":"50rpx","fontSize":"28rpx","minWidth":"160rpx","height":"50rpx"}' class="label">出生日期：</view>
				<picker  :style='{"width":"100%","margin":"0","display":"flex","height":"auto"}' v-if="tableName=='guke'" mode="date" :value="ruleForm.chushengriqi" @change="gukechushengriqiDateChange">
					<view>
						<view :style='{"padding":"0 20rpx","lineHeight":"80rpx","fontSize":"28rpx","color":"#74868B","flex":"1"}' class="uni-input">{{ruleForm.chushengriqi?ruleForm.chushengriqi:"请选择出生日期"}}</view>
					</view>
				</picker>
			</view>
			<view :style='{"border":"none","width":"100%","margin":"0 0 40rpx 0","alignItems":"center","display":"flex","height":"auto"}' v-if="tableName=='guke'" class="uni-form-item uni-column">
				<view :style='{"padding":"0","color":"#000000","borderRadius":"0","textAlign":"center","background":"none","flex":"none","width":"auto","lineHeight":"50rpx","fontSize":"28rpx","minWidth":"160rpx","height":"50rpx"}' class="label">年龄：</view>
				<input :style='{"padding":"0px 24rpx","margin":"0px","borderColor":"#DFDFDF","color":"#000","borderRadius":"0","flex":"1","background":"none","borderWidth":"0 0 2rpx 0","fontSize":"28rpx","borderStyle":"solid","height":"80rpx"}'  v-model="ruleForm.nianling"  type="text"  class="uni-input" name="" placeholder="年龄" />
			</view>
			<view :style='{"border":"none","width":"100%","margin":"0 0 40rpx 0","alignItems":"center","display":"flex","height":"auto"}' v-if="tableName=='guke'" class="uni-form-item uni-column">
				<view :style='{"padding":"0","color":"#000000","borderRadius":"0","textAlign":"center","background":"none","flex":"none","width":"auto","lineHeight":"50rpx","fontSize":"28rpx","minWidth":"160rpx","height":"50rpx"}' class="label">学历：</view>
				<input :style='{"padding":"0px 24rpx","margin":"0px","borderColor":"#DFDFDF","color":"#000","borderRadius":"0","flex":"1","background":"none","borderWidth":"0 0 2rpx 0","fontSize":"28rpx","borderStyle":"solid","height":"80rpx"}'  v-model="ruleForm.xueli"  type="text"  class="uni-input" name="" placeholder="学历" />
			</view>
			<view :style='{"border":"none","width":"100%","margin":"0 0 40rpx 0","alignItems":"center","display":"flex","height":"auto"}' v-if="tableName=='guke'" class="uni-form-item uni-column">
				<view :style='{"padding":"0","color":"#000000","borderRadius":"0","textAlign":"center","background":"none","flex":"none","width":"auto","lineHeight":"50rpx","fontSize":"28rpx","minWidth":"160rpx","height":"50rpx"}' class="label">住址：</view>
				<input :style='{"padding":"0px 24rpx","margin":"0px","borderColor":"#DFDFDF","color":"#000","borderRadius":"0","flex":"1","background":"none","borderWidth":"0 0 2rpx 0","fontSize":"28rpx","borderStyle":"solid","height":"80rpx"}'  v-model="ruleForm.zhuzhi"  type="text"  class="uni-input" name="" placeholder="住址" />
			</view>
			<view :style='{"border":"none","width":"100%","margin":"0 0 40rpx 0","alignItems":"center","display":"flex","height":"auto"}' v-if="tableName=='yuangong'" class="uni-form-item uni-column">
				<view :style='{"padding":"0","color":"#000000","borderRadius":"0","textAlign":"center","background":"none","flex":"none","width":"auto","lineHeight":"50rpx","fontSize":"28rpx","minWidth":"160rpx","height":"50rpx"}' class="label">工号：</view>
				<input :style='{"padding":"0px 24rpx","margin":"0px","borderColor":"#DFDFDF","color":"#000","borderRadius":"0","flex":"1","background":"none","borderWidth":"0 0 2rpx 0","fontSize":"28rpx","borderStyle":"solid","height":"80rpx"}'  v-model="ruleForm.gonghao"  type="text"  class="uni-input" name="" placeholder="工号" />
			</view>
			<view :style='{"border":"none","width":"100%","margin":"0 0 40rpx 0","alignItems":"center","display":"flex","height":"auto"}' v-if="tableName=='yuangong'" class="uni-form-item uni-column">
				<view :style='{"padding":"0","color":"#000000","borderRadius":"0","textAlign":"center","background":"none","flex":"none","width":"auto","lineHeight":"50rpx","fontSize":"28rpx","minWidth":"160rpx","height":"50rpx"}' class="label">密码：</view>
				<input :style='{"padding":"0px 24rpx","margin":"0px","borderColor":"#DFDFDF","color":"#000","borderRadius":"0","flex":"1","background":"none","borderWidth":"0 0 2rpx 0","fontSize":"28rpx","borderStyle":"solid","height":"80rpx"}'  v-model="ruleForm.mima" type="password"  class="uni-input" name="" placeholder="密码" />
			</view>
			<view :style='{"border":"none","width":"100%","margin":"0 0 40rpx 0","alignItems":"center","display":"flex","height":"auto"}' v-if="tableName=='yuangong'" class="uni-form-item uni-column">
				<view :style='{"padding":"0","color":"#000000","borderRadius":"0","textAlign":"center","background":"none","flex":"none","width":"auto","lineHeight":"50rpx","fontSize":"28rpx","minWidth":"160rpx","height":"50rpx"}' class="label">确认密码：</view>
				<input :style='{"padding":"0px 24rpx","margin":"0px","borderColor":"#DFDFDF","color":"#000","borderRadius":"0","flex":"1","background":"none","borderWidth":"0 0 2rpx 0","fontSize":"28rpx","borderStyle":"solid","height":"80rpx"}' v-model="ruleForm.mima2" type="password" class="uni-input" name="" placeholder="确认密码" />
			</view>
			<view :style='{"border":"none","width":"100%","margin":"0 0 40rpx 0","alignItems":"center","display":"flex","height":"auto"}' v-if="tableName=='yuangong'" class="uni-form-item uni-column">
				<view :style='{"padding":"0","color":"#000000","borderRadius":"0","textAlign":"center","background":"none","flex":"none","width":"auto","lineHeight":"50rpx","fontSize":"28rpx","minWidth":"160rpx","height":"50rpx"}' class="label">员工姓名：</view>
				<input :style='{"padding":"0px 24rpx","margin":"0px","borderColor":"#DFDFDF","color":"#000","borderRadius":"0","flex":"1","background":"none","borderWidth":"0 0 2rpx 0","fontSize":"28rpx","borderStyle":"solid","height":"80rpx"}'  v-model="ruleForm.yuangongxingming"  type="text"  class="uni-input" name="" placeholder="员工姓名" />
			</view>
            <view :style='{"border":"none","width":"100%","margin":"0 0 40rpx 0","alignItems":"center","display":"flex","height":"auto"}' v-if="tableName=='yuangong'" class="">
				<view :style='{"padding":"0","color":"#000000","borderRadius":"0","textAlign":"center","background":"none","flex":"none","width":"auto","lineHeight":"50rpx","fontSize":"28rpx","minWidth":"160rpx","height":"50rpx"}' class="label">性别：</view>
				<picker :style='{"width":"100%","margin":"0","display":"flex","height":"auto"}' v-if="tableName=='yuangong'"  @change="yuangongxingbieChange" :value="yuangongxingbieIndex" :range="yuangongxingbieOptions">
					<view>
						<view :style='{"padding":"0 20rpx","lineHeight":"80rpx","fontSize":"28rpx","color":"#74868B","flex":"1"}' class="uni-input">{{ruleForm.xingbie?ruleForm.xingbie:"请选择性别"}}</view>
					</view>
				</picker>
			</view>
            <view :style='{"border":"none","width":"100%","margin":"0 0 40rpx 0","alignItems":"center","display":"flex","height":"auto"}' v-if="tableName=='yuangong'" @tap="yuangongtouxiangTap" class="">
				<view :style='{"padding":"0","color":"#000000","borderRadius":"0","textAlign":"center","background":"none","flex":"none","width":"auto","lineHeight":"50rpx","fontSize":"28rpx","minWidth":"160rpx","height":"50rpx"}' class="label">头像：</view>
				<image :style='{"width":"120rpx","borderRadius":"16rpx","display":"block","height":"120rpx"}' v-if="ruleForm.touxiang" class="avator" :src="baseUrl+ruleForm.touxiang" mode=""></image>
                <image :style='{"width":"120rpx","borderRadius":"16rpx","display":"block","height":"120rpx"}' v-else class="avator" src="../../static/gen/upload.png" mode=""></image>
            </view>
			<view :style='{"border":"none","width":"100%","margin":"0 0 40rpx 0","alignItems":"center","display":"flex","height":"auto"}' v-if="tableName=='yuangong'" class="uni-form-item uni-column">
				<view :style='{"padding":"0","color":"#000000","borderRadius":"0","textAlign":"center","background":"none","flex":"none","width":"auto","lineHeight":"50rpx","fontSize":"28rpx","minWidth":"160rpx","height":"50rpx"}' class="label">身份证：</view>
				<input :style='{"padding":"0px 24rpx","margin":"0px","borderColor":"#DFDFDF","color":"#000","borderRadius":"0","flex":"1","background":"none","borderWidth":"0 0 2rpx 0","fontSize":"28rpx","borderStyle":"solid","height":"80rpx"}'  v-model="ruleForm.shenfenzheng"  type="text"  class="uni-input" name="" placeholder="身份证" />
			</view>
			<view :style='{"border":"none","width":"100%","margin":"0 0 40rpx 0","alignItems":"center","display":"flex","height":"auto"}' v-if="tableName=='yuangong'" class="uni-form-item uni-column">
				<view :style='{"padding":"0","color":"#000000","borderRadius":"0","textAlign":"center","background":"none","flex":"none","width":"auto","lineHeight":"50rpx","fontSize":"28rpx","minWidth":"160rpx","height":"50rpx"}' class="label">电话：</view>
				<input :style='{"padding":"0px 24rpx","margin":"0px","borderColor":"#DFDFDF","color":"#000","borderRadius":"0","flex":"1","background":"none","borderWidth":"0 0 2rpx 0","fontSize":"28rpx","borderStyle":"solid","height":"80rpx"}'  v-model="ruleForm.dianhua"  type="text"  class="uni-input" name="" placeholder="电话" />
			</view>
			<view :style='{"border":"none","width":"100%","margin":"0 0 40rpx 0","alignItems":"center","display":"flex","height":"auto"}' v-if="tableName=='yuangong'" class="uni-form-item uni-column">
				<view :style='{"padding":"0","color":"#000000","borderRadius":"0","textAlign":"center","background":"none","flex":"none","width":"auto","lineHeight":"50rpx","fontSize":"28rpx","minWidth":"160rpx","height":"50rpx"}' class="label">民族：</view>
				<input :style='{"padding":"0px 24rpx","margin":"0px","borderColor":"#DFDFDF","color":"#000","borderRadius":"0","flex":"1","background":"none","borderWidth":"0 0 2rpx 0","fontSize":"28rpx","borderStyle":"solid","height":"80rpx"}'  v-model="ruleForm.minzu"  type="text"  class="uni-input" name="" placeholder="民族" />
			</view>
			<view :style='{"border":"none","width":"100%","margin":"0 0 40rpx 0","alignItems":"center","display":"flex","height":"auto"}' v-if="tableName=='yuangong'" class="uni-form-item uni-column">
				<view :style='{"padding":"0","color":"#000000","borderRadius":"0","textAlign":"center","background":"none","flex":"none","width":"auto","lineHeight":"50rpx","fontSize":"28rpx","minWidth":"160rpx","height":"50rpx"}' class="label">籍贯：</view>
				<input :style='{"padding":"0px 24rpx","margin":"0px","borderColor":"#DFDFDF","color":"#000","borderRadius":"0","flex":"1","background":"none","borderWidth":"0 0 2rpx 0","fontSize":"28rpx","borderStyle":"solid","height":"80rpx"}'  v-model="ruleForm.jiguan"  type="text"  class="uni-input" name="" placeholder="籍贯" />
			</view>
            <view :style='{"border":"none","width":"100%","margin":"0 0 40rpx 0","alignItems":"center","display":"flex","height":"auto"}' v-if="tableName=='yuangong'" class="">
				<view :style='{"padding":"0","color":"#000000","borderRadius":"0","textAlign":"center","background":"none","flex":"none","width":"auto","lineHeight":"50rpx","fontSize":"28rpx","minWidth":"160rpx","height":"50rpx"}' class="label">出生日月：</view>
				<picker  :style='{"width":"100%","margin":"0","display":"flex","height":"auto"}' v-if="tableName=='yuangong'" mode="date" :value="ruleForm.chushengriyue" @change="yuangongchushengriyueDateChange">
					<view>
						<view :style='{"padding":"0 20rpx","lineHeight":"80rpx","fontSize":"28rpx","color":"#74868B","flex":"1"}' class="uni-input">{{ruleForm.chushengriyue?ruleForm.chushengriyue:"请选择出生日月"}}</view>
					</view>
				</picker>
			</view>
			<view :style='{"border":"none","width":"100%","margin":"0 0 40rpx 0","alignItems":"center","display":"flex","height":"auto"}' v-if="tableName=='yuangong'" class="uni-form-item uni-column">
				<view :style='{"padding":"0","color":"#000000","borderRadius":"0","textAlign":"center","background":"none","flex":"none","width":"auto","lineHeight":"50rpx","fontSize":"28rpx","minWidth":"160rpx","height":"50rpx"}' class="label">年龄：</view>
				<input :style='{"padding":"0px 24rpx","margin":"0px","borderColor":"#DFDFDF","color":"#000","borderRadius":"0","flex":"1","background":"none","borderWidth":"0 0 2rpx 0","fontSize":"28rpx","borderStyle":"solid","height":"80rpx"}'  v-model="ruleForm.nianling"  type="text"  class="uni-input" name="" placeholder="年龄" />
			</view>
			<view :style='{"border":"none","width":"100%","margin":"0 0 40rpx 0","alignItems":"center","display":"flex","height":"auto"}' v-if="tableName=='yuangong'" class="uni-form-item uni-column">
				<view :style='{"padding":"0","color":"#000000","borderRadius":"0","textAlign":"center","background":"none","flex":"none","width":"auto","lineHeight":"50rpx","fontSize":"28rpx","minWidth":"160rpx","height":"50rpx"}' class="label">学历：</view>
				<input :style='{"padding":"0px 24rpx","margin":"0px","borderColor":"#DFDFDF","color":"#000","borderRadius":"0","flex":"1","background":"none","borderWidth":"0 0 2rpx 0","fontSize":"28rpx","borderStyle":"solid","height":"80rpx"}'  v-model="ruleForm.xueli"  type="text"  class="uni-input" name="" placeholder="学历" />
			</view>
			<view :style='{"border":"none","width":"100%","margin":"0 0 40rpx 0","alignItems":"center","display":"flex","height":"auto"}' v-if="tableName=='yuangong'" class="uni-form-item uni-column">
				<view :style='{"padding":"0","color":"#000000","borderRadius":"0","textAlign":"center","background":"none","flex":"none","width":"auto","lineHeight":"50rpx","fontSize":"28rpx","minWidth":"160rpx","height":"50rpx"}' class="label">住址：</view>
				<input :style='{"padding":"0px 24rpx","margin":"0px","borderColor":"#DFDFDF","color":"#000","borderRadius":"0","flex":"1","background":"none","borderWidth":"0 0 2rpx 0","fontSize":"28rpx","borderStyle":"solid","height":"80rpx"}'  v-model="ruleForm.zhuzhi"  type="text"  class="uni-input" name="" placeholder="住址" />
			</view>
			<button :style='{"border":"0","padding":"0px","margin":"0 0 24rpx 0","color":"#000","borderRadius":"8rpx","background":"linear-gradient( 90deg, #EFE0BE 0%, #E4EFBE 100%)","width":"100%","lineHeight":"80rpx","fontSize":"32rpx","fontWeight":"600","height":"80rpx"}' class="btn-submit" @tap="register" type="primary">注册</button>
			
			<view class="idea1" :style='{"width":"100%","background":"red","display":"none","height":"80rpx"}'>idea1</view>
			<view class="idea2" :style='{"width":"100%","background":"red","display":"none","height":"80rpx"}'>idea2</view>
			<view class="idea3" :style='{"width":"100%","background":"red","display":"none","height":"80rpx"}'>idea3</view>
		</view>
	</view>
</view>
</template>

<script>
    import multipleSelect from "@/components/momo-multipleSelect/momo-multipleSelect";
	export default {
		data() {
			return {
                gukexingbieOptions: [],
                gukexingbieIndex: 0,
                yuangongxingbieOptions: [],
                yuangongxingbieIndex: 0,
				ruleForm: {
                zhanghao: '',
                mima: '',
                xingming: '',
                xingbie: '',
                shouji: '',
                touxiang: '',
                shenfenzheng: '',
                minzu: '',
                jiguan: '',
                chushengriqi: '',
                nianling: '',
                xueli: '',
                zhuzhi: '',
                xuqiu: '',
                gonghao: '',
                mima: '',
                yuangongxingming: '',
                xingbie: '',
                touxiang: '',
                shenfenzheng: '',
                dianhua: '',
                minzu: '',
                jiguan: '',
                chushengriyue: '',
                nianling: '',
                xueli: '',
                zhuzhi: '',
                tezhang: '',
                ziwopingjia: '',
                gongzuoanpai: '',
                gongzixinxi: '',
                jifen: '0',
				},
				tableName:""
			}
		},
        components: {
            multipleSelect
        },
        computed: {
            baseUrl() {
                return this.$base.url;
            },
        },
		async onLoad() {
			let res = [];
			let table = uni.getStorageSync("loginTable");
            this.tableName = table;

                        // 自定义下拉框值
			if(this.tableName=='guke'){
                this.gukexingbieOptions = "男,女".split(',');
				this.ruleForm.xingbie=this.gukexingbieOptions[0]
			}
                        // 自定义下拉框值
			if(this.tableName=='yuangong'){
                this.yuangongxingbieOptions = "男,女".split(',');
				this.ruleForm.xingbie=this.yuangongxingbieOptions[0]
			}
            if(`yuangong` == this.tableName){
                this.ruleForm.jifen='0'
            }
			
			this.styleChange()
		},
		methods: {

            // 下拉变化
            gukexingbieChange(e) {
                    this.gukexingbieIndex = e.target.value
                    this.ruleForm.xingbie = this.gukexingbieOptions[this.gukexingbieIndex]
            },
            guketouxiangTap() {
                let _this = this;
                this.$api.upload(function(res) {
                    _this.ruleForm.touxiang = 'upload/' + res.file;
					_this.$forceUpdate();
                });
            },
            gukechushengriqiDateChange(e) {
                this.ruleForm.chushengriqi = e.target.value;
                this.$forceUpdate();
            },
            // 下拉变化
            yuangongxingbieChange(e) {
                    this.yuangongxingbieIndex = e.target.value
                    this.ruleForm.xingbie = this.yuangongxingbieOptions[this.yuangongxingbieIndex]
            },
            yuangongtouxiangTap() {
                let _this = this;
                this.$api.upload(function(res) {
                    _this.ruleForm.touxiang = 'upload/' + res.file;
					_this.$forceUpdate();
                });
            },
            yuangongchushengriyueDateChange(e) {
                this.ruleForm.chushengriyue = e.target.value;
                this.$forceUpdate();
            },

            toggleTab(str) {
                this.$refs[str].show();
            },

			styleChange() {
				this.$nextTick(()=>{
					// document.querySelectorAll('.uni-input .uni-input-input').forEach(el=>{
					//   el.style.backgroundColor = this.registerFrom.content.input.backgroundColor
					// })
				})
			},
			// 获取uuid
			getUUID () {
				return new Date().getTime();
			},
			// 注册
			async register() {
				
				if((!this.ruleForm.zhanghao) && `guke` == this.tableName){
					this.$utils.msg(`账号不能为空`);
					return
				}
				if((!this.ruleForm.mima) && `guke` == this.tableName){
					this.$utils.msg(`密码不能为空`);
					return
				}
                if(`guke` == this.tableName && (this.ruleForm.mima!=this.ruleForm.mima2)){
                    this.$utils.msg(`两次密码输入不一致`);
                    return
                }
				if((!this.ruleForm.xingming) && `guke` == this.tableName){
					this.$utils.msg(`姓名不能为空`);
					return
				}
				if(`guke` == this.tableName && this.ruleForm.shouji&&(!this.$validate.isMobile(this.ruleForm.shouji))){
					this.$utils.msg(`手机应输入手机格式`);
					return
				}
				if(`guke` == this.tableName && this.ruleForm.shenfenzheng&&(!this.$validate.checkIdCard(this.ruleForm.shenfenzheng))){
					this.$utils.msg(`身份证应输入身份证格式`);
					return
				}
				if(`guke` == this.tableName && this.ruleForm.nianling&&(!this.$validate.isIntNumer(this.ruleForm.nianling))){
					this.$utils.msg(`年龄应输入整数`);
					return
				}
				if((!this.ruleForm.gonghao) && `yuangong` == this.tableName){
					this.$utils.msg(`工号不能为空`);
					return
				}
				if((!this.ruleForm.mima) && `yuangong` == this.tableName){
					this.$utils.msg(`密码不能为空`);
					return
				}
                if(`yuangong` == this.tableName && (this.ruleForm.mima!=this.ruleForm.mima2)){
                    this.$utils.msg(`两次密码输入不一致`);
                    return
                }
				if((!this.ruleForm.yuangongxingming) && `yuangong` == this.tableName){
					this.$utils.msg(`员工姓名不能为空`);
					return
				}
				if(`yuangong` == this.tableName && this.ruleForm.shenfenzheng&&(!this.$validate.checkIdCard(this.ruleForm.shenfenzheng))){
					this.$utils.msg(`身份证应输入身份证格式`);
					return
				}
				if(`yuangong` == this.tableName && this.ruleForm.dianhua&&(!this.$validate.isMobile(this.ruleForm.dianhua))){
					this.$utils.msg(`电话应输入手机格式`);
					return
				}
				if(`yuangong` == this.tableName && this.ruleForm.nianling&&(!this.$validate.isIntNumer(this.ruleForm.nianling))){
					this.$utils.msg(`年龄应输入整数`);
					return
				}
				if(`yuangong` == this.tableName && this.ruleForm.jifen&&(!this.$validate.isIntNumer(this.ruleForm.jifen))){
					this.$utils.msg(`积分应输入整数`);
					return
				}
				await this.$api.register(`${this.tableName}`, this.ruleForm);
				this.$utils.msgBack('注册成功');;
			}
		}
	}
</script>

<style lang="scss" scoped>
	.content {
		min-height: calc(100vh - 44px);
		box-sizing: border-box;
	}
</style>
