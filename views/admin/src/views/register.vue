<template>
	<div>
		<div class="register-container">
			<el-form v-if="pageFlag=='register'" ref="ruleForm" class="rgs-form animate__animated animate__backInDown" :model="ruleForm" :rules="rules">
				<div class="rgs-form2">
					<div class="title">骄阳房屋租赁公司业务管理系统</div>
					<el-form-item class="list-item" v-if="tableName=='guke'">
						<div class="lable" :class="changeRules('zhanghao')?'required':''">账号：</div>
						<el-input  v-model="ruleForm.zhanghao"  autocomplete="off" placeholder="账号"  type="text"  />
					</el-form-item>
					<el-form-item class="list-item" v-if="tableName=='guke'">
						<div class="lable" :class="changeRules('mima')?'required':''">密码：</div>
						<el-input  v-model="ruleForm.mima"  autocomplete="off" placeholder="密码"  type="password"  />
					</el-form-item>
					<el-form-item class="list-item" v-if="tableName=='guke'">
						<div class="lable" :class="changeRules('mima')?'required':''">确认密码：</div>
						<el-input  v-model="ruleForm.mima2" autocomplete="off" placeholder="确认密码" type="password" />
					</el-form-item>
					<el-form-item class="list-item" v-if="tableName=='guke'">
						<div class="lable" :class="changeRules('xingming')?'required':''">姓名：</div>
						<el-input  v-model="ruleForm.xingming"  autocomplete="off" placeholder="姓名"  type="text"  />
					</el-form-item>
					<el-form-item class="list-item" v-if="tableName=='guke'">
						<div class="lable" :class="changeRules('xingbie')?'required':''">性别：</div>
						<el-select v-model="ruleForm.xingbie" placeholder="请选择性别" >
							<el-option
								v-for="(item,index) in gukexingbieOptions"
								v-bind:key="index"
								:label="item"
								:value="item">
							</el-option>
						</el-select>
					</el-form-item>
					<el-form-item class="list-item" v-if="tableName=='guke'">
						<div class="lable" :class="changeRules('shouji')?'required':''">手机：</div>
						<el-input  v-model="ruleForm.shouji"  autocomplete="off" placeholder="手机"  type="text"  />
					</el-form-item>
					<el-form-item class="list-item" v-if="tableName=='guke'">
						<div class="lable" :class="changeRules('touxiang')?'required':''">头像：</div>
						<file-upload
							tip="点击上传头像"
							action="file/upload"
							:limit="3"
							:multiple="true"
							:fileUrls="ruleForm.touxiang?ruleForm.touxiang:''"
							@change="guketouxiangUploadChange"
						></file-upload>
					</el-form-item>
					<el-form-item class="list-item" v-if="tableName=='guke'">
						<div class="lable" :class="changeRules('shenfenzheng')?'required':''">身份证：</div>
						<el-input  v-model="ruleForm.shenfenzheng"  autocomplete="off" placeholder="身份证"  type="text"  />
					</el-form-item>
					<el-form-item class="list-item" v-if="tableName=='guke'">
						<div class="lable" :class="changeRules('minzu')?'required':''">民族：</div>
						<el-input  v-model="ruleForm.minzu"  autocomplete="off" placeholder="民族"  type="text"  />
					</el-form-item>
					<el-form-item class="list-item" v-if="tableName=='guke'">
						<div class="lable" :class="changeRules('jiguan')?'required':''">籍贯：</div>
						<el-input  v-model="ruleForm.jiguan"  autocomplete="off" placeholder="籍贯"  type="text"  />
					</el-form-item>
					<el-form-item class="list-item" v-if="tableName=='guke'">
						<div class="lable" :class="changeRules('chushengriqi')?'required':''">出生日期：</div>
						<el-date-picker
							format="yyyy 年 MM 月 dd 日"
							value-format="yyyy-MM-dd"
							v-model="ruleForm.chushengriqi"
							type="date"
							placeholder="出生日期"
						></el-date-picker>
					</el-form-item>
					<el-form-item class="list-item" v-if="tableName=='guke'">
						<div class="lable" :class="changeRules('nianling')?'required':''">年龄：</div>
						<el-input  v-model.number="ruleForm.nianling"  autocomplete="off" placeholder="年龄"  type="text"  />
					</el-form-item>
					<el-form-item class="list-item" v-if="tableName=='guke'">
						<div class="lable" :class="changeRules('xueli')?'required':''">学历：</div>
						<el-input  v-model="ruleForm.xueli"  autocomplete="off" placeholder="学历"  type="text"  />
					</el-form-item>
					<el-form-item class="list-item" v-if="tableName=='guke'">
						<div class="lable" :class="changeRules('zhuzhi')?'required':''">住址：</div>
						<el-input  v-model="ruleForm.zhuzhi"  autocomplete="off" placeholder="住址"  type="text"  />
					</el-form-item>
					<el-form-item class="list-item" v-if="tableName=='yuangong'">
						<div class="lable" :class="changeRules('gonghao')?'required':''">工号：</div>
						<el-input  v-model="ruleForm.gonghao"  autocomplete="off" placeholder="工号"  type="text"  />
					</el-form-item>
					<el-form-item class="list-item" v-if="tableName=='yuangong'">
						<div class="lable" :class="changeRules('mima')?'required':''">密码：</div>
						<el-input  v-model="ruleForm.mima"  autocomplete="off" placeholder="密码"  type="password"  />
					</el-form-item>
					<el-form-item class="list-item" v-if="tableName=='yuangong'">
						<div class="lable" :class="changeRules('mima')?'required':''">确认密码：</div>
						<el-input  v-model="ruleForm.mima2" autocomplete="off" placeholder="确认密码" type="password" />
					</el-form-item>
					<el-form-item class="list-item" v-if="tableName=='yuangong'">
						<div class="lable" :class="changeRules('yuangongxingming')?'required':''">员工姓名：</div>
						<el-input  v-model="ruleForm.yuangongxingming"  autocomplete="off" placeholder="员工姓名"  type="text"  />
					</el-form-item>
					<el-form-item class="list-item" v-if="tableName=='yuangong'">
						<div class="lable" :class="changeRules('xingbie')?'required':''">性别：</div>
						<el-select v-model="ruleForm.xingbie" placeholder="请选择性别" >
							<el-option
								v-for="(item,index) in yuangongxingbieOptions"
								v-bind:key="index"
								:label="item"
								:value="item">
							</el-option>
						</el-select>
					</el-form-item>
					<el-form-item class="list-item" v-if="tableName=='yuangong'">
						<div class="lable" :class="changeRules('touxiang')?'required':''">头像：</div>
						<file-upload
							tip="点击上传头像"
							action="file/upload"
							:limit="3"
							:multiple="true"
							:fileUrls="ruleForm.touxiang?ruleForm.touxiang:''"
							@change="yuangongtouxiangUploadChange"
						></file-upload>
					</el-form-item>
					<el-form-item class="list-item" v-if="tableName=='yuangong'">
						<div class="lable" :class="changeRules('shenfenzheng')?'required':''">身份证：</div>
						<el-input  v-model="ruleForm.shenfenzheng"  autocomplete="off" placeholder="身份证"  type="text"  />
					</el-form-item>
					<el-form-item class="list-item" v-if="tableName=='yuangong'">
						<div class="lable" :class="changeRules('dianhua')?'required':''">电话：</div>
						<el-input  v-model="ruleForm.dianhua"  autocomplete="off" placeholder="电话"  type="text"  />
					</el-form-item>
					<el-form-item class="list-item" v-if="tableName=='yuangong'">
						<div class="lable" :class="changeRules('minzu')?'required':''">民族：</div>
						<el-input  v-model="ruleForm.minzu"  autocomplete="off" placeholder="民族"  type="text"  />
					</el-form-item>
					<el-form-item class="list-item" v-if="tableName=='yuangong'">
						<div class="lable" :class="changeRules('jiguan')?'required':''">籍贯：</div>
						<el-input  v-model="ruleForm.jiguan"  autocomplete="off" placeholder="籍贯"  type="text"  />
					</el-form-item>
					<el-form-item class="list-item" v-if="tableName=='yuangong'">
						<div class="lable" :class="changeRules('chushengriyue')?'required':''">出生日月：</div>
						<el-date-picker
							format="yyyy 年 MM 月 dd 日"
							value-format="yyyy-MM-dd"
							v-model="ruleForm.chushengriyue"
							type="date"
							placeholder="出生日月"
						></el-date-picker>
					</el-form-item>
					<el-form-item class="list-item" v-if="tableName=='yuangong'">
						<div class="lable" :class="changeRules('nianling')?'required':''">年龄：</div>
						<el-input  v-model.number="ruleForm.nianling"  autocomplete="off" placeholder="年龄"  type="text"  />
					</el-form-item>
					<el-form-item class="list-item" v-if="tableName=='yuangong'">
						<div class="lable" :class="changeRules('xueli')?'required':''">学历：</div>
						<el-input  v-model="ruleForm.xueli"  autocomplete="off" placeholder="学历"  type="text"  />
					</el-form-item>
					<el-form-item class="list-item" v-if="tableName=='yuangong'">
						<div class="lable" :class="changeRules('zhuzhi')?'required':''">住址：</div>
						<el-input  v-model="ruleForm.zhuzhi"  autocomplete="off" placeholder="住址"  type="text"  />
					</el-form-item>
					<div class="register-btn">
						<div class="register-btn1">
							<button type="button" class="r-btn" @click="login()">注册</button>
						</div>
						<div class="register-btn2">
							<div class="r-login" @click="close()">已有账号，直接登录</div>
						</div>
					</div>
				</div>
				<div class="idea-box2">输入您的账号和密码以注册帐户</div>
			</el-form>
		</div>
	</div>
</template>

<script>
	import 'animate.css'
export default {
	data() {
		return {
			ruleForm: {
			},
			forgetForm: {},
            pageFlag : '',
			tableName:"",
			rules: {},
            gukexingbieOptions: [],
            yuangongxingbieOptions: [],
		};
	},
	mounted(){
		this.pageFlag = this.$route.query.pageFlag
		if(this.$route.query.pageFlag=='register'){
			
			let table = this.$storage.get("loginTable");
			this.tableName = table;
			if(this.tableName=='guke'){
				this.ruleForm = {
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
				}
			}
			if(this.tableName=='yuangong'){
				this.ruleForm = {
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
				}
			}
			if ('guke' == this.tableName) {
				this.rules.zhanghao = [{ required: true, message: '请输入账号', trigger: 'blur' }]
			}
			if ('guke' == this.tableName) {
				this.rules.mima = [{ required: true, message: '请输入密码', trigger: 'blur' }]
			}
			if ('guke' == this.tableName) {
				this.rules.xingming = [{ required: true, message: '请输入姓名', trigger: 'blur' }]
			}
			if ('yuangong' == this.tableName) {
				this.rules.gonghao = [{ required: true, message: '请输入工号', trigger: 'blur' }]
			}
			if ('yuangong' == this.tableName) {
				this.rules.mima = [{ required: true, message: '请输入密码', trigger: 'blur' }]
			}
			if ('yuangong' == this.tableName) {
				this.rules.yuangongxingming = [{ required: true, message: '请输入员工姓名', trigger: 'blur' }]
			}
			this.gukexingbieOptions = "男,女".split(',')
			this.yuangongxingbieOptions = "男,女".split(',')
		}
	},
	created() {
	},
	destroyed() {
		  	},
	methods: {
		changeRules(name){
			if(this.rules[name]){
				return true
			}
			return false
		},
		// 获取uuid
		getUUID () {
			return new Date().getTime();
		},
		close(){
			this.$router.push({ path: "/login" });
		},
        guketouxiangUploadChange(fileUrls) {
            this.ruleForm.touxiang = fileUrls;
        },
        yuangongtouxiangUploadChange(fileUrls) {
            this.ruleForm.touxiang = fileUrls;
        },

        // 多级联动参数


		// 注册
		login() {
			var url=this.tableName+"/register";
			if((!this.ruleForm.zhanghao) && `guke` == this.tableName){
				this.$message.error(`账号不能为空`);
				return
			}
			if((!this.ruleForm.mima) && `guke` == this.tableName){
				this.$message.error(`密码不能为空`);
				return
			}
			if((this.ruleForm.mima!=this.ruleForm.mima2) && `guke` == this.tableName){
				this.$message.error(`两次密码输入不一致`);
				return
			}
			if((!this.ruleForm.xingming) && `guke` == this.tableName){
				this.$message.error(`姓名不能为空`);
				return
			}
			if(`guke` == this.tableName && this.ruleForm.shouji &&(!this.$validate.isMobile(this.ruleForm.shouji))){
				this.$message.error(`手机应输入手机格式`);
				return
			}
            if(this.ruleForm.touxiang!=null) {
                this.ruleForm.touxiang = this.ruleForm.touxiang.replace(new RegExp(this.$base.url,"g"),"");
            }
			if(`guke` == this.tableName && this.ruleForm.shenfenzheng &&(!this.$validate.checkIdCard(this.ruleForm.shenfenzheng))){
				this.$message.error(`身份证应输入身份证格式`);
				return
			}
			if(`guke` == this.tableName && this.ruleForm.nianling &&(!this.$validate.isIntNumer(this.ruleForm.nianling))){
				this.$message.error(`年龄应输入整数`);
				return
			}
			if((!this.ruleForm.gonghao) && `yuangong` == this.tableName){
				this.$message.error(`工号不能为空`);
				return
			}
			if((!this.ruleForm.mima) && `yuangong` == this.tableName){
				this.$message.error(`密码不能为空`);
				return
			}
			if((this.ruleForm.mima!=this.ruleForm.mima2) && `yuangong` == this.tableName){
				this.$message.error(`两次密码输入不一致`);
				return
			}
			if((!this.ruleForm.yuangongxingming) && `yuangong` == this.tableName){
				this.$message.error(`员工姓名不能为空`);
				return
			}
            if(this.ruleForm.touxiang!=null) {
                this.ruleForm.touxiang = this.ruleForm.touxiang.replace(new RegExp(this.$base.url,"g"),"");
            }
			if(`yuangong` == this.tableName && this.ruleForm.shenfenzheng &&(!this.$validate.checkIdCard(this.ruleForm.shenfenzheng))){
				this.$message.error(`身份证应输入身份证格式`);
				return
			}
			if(`yuangong` == this.tableName && this.ruleForm.dianhua &&(!this.$validate.isMobile(this.ruleForm.dianhua))){
				this.$message.error(`电话应输入手机格式`);
				return
			}
			if(`yuangong` == this.tableName && this.ruleForm.nianling &&(!this.$validate.isIntNumer(this.ruleForm.nianling))){
				this.$message.error(`年龄应输入整数`);
				return
			}
			this.$http({
				url: url,
				method: "post",
				data:this.ruleForm
			}).then(({ data }) => {
				if (data && data.code === 0) {
					this.$message({
						message: "注册成功",
						type: "success",
						duration: 1500,
						onClose: () => {
							this.$router.replace({ path: "/login" });
						}
					});
				} else {
					this.$message.error(data.msg);
				}
			});
		}
	}
};
</script>

<style lang="scss" scoped>
.register-container {
	position: relative;
	background-repeat: no-repeat;
	background-size: 100% 100%;
	display: flex;
	width: 100%;
	min-height: 100vh;
	justify-content: center;
	align-items: center;
	background-image:  url(http://codegen.caihongy.cn/20240926/a12a8b82de814927adbd5783c98082da.jpg);
	background-position: center center;
	position: relative;
	.rgs-form {
		.rgs-form2 {
		background: #fff;
		width: 75%;
		}
		border-radius: 0;
		padding: 0px 69px 40px 25%;
		box-shadow: inset 0px 0px 0px 0px #000;
		margin: auto;
		z-index: 1000;
		flex-direction: column;
		background: url("http://codegen.caihongy.cn/20240926/ee2e080ca41941cf8b9064ea8d1df11d.png") left center /  48% 130% no-repeat, #fff;
		display: flex;
		width: 1200px;
		align-items: flex-end;
		height: auto;
		.title {
			padding: 0 0px;
			margin: 30px 30px 30px -120px;
			color: #000000;
			background: none;
			font-weight: 500;
			width: calc(100% + 160px);
			font-size: 22px;
			font-family: Source Han Sans-Medium;
			line-height: 20px;
			text-align: center;
		}
		.list-item {
			padding: 0 0 0 0px;
			margin: 0 0 15px 120px;
			width: calc(100% - 120px);
			position: relative;
			height: auto;
			/deep/ .el-form-item__content {
				display: block;
			}
			.lable {
				padding: 0 10px 0 0;
				color: #000;
				left: -120px;
				letter-spacing: 1px;
				width: 120px;
				font-size: 16px;
				border-color: #000000;
				border-width: 0 0 0px 0;
				position: absolute!important;
				border-style: solid;
				text-align: right;
				height: 46px;
			}
			.el-input {
				width: 100%;
			}
			.el-input /deep/ .el-input__inner {
				border-radius: 0;
				padding: 0 20px;
				color: #666;
				background: #fff;
				flex: 1;
				font-size: 16px;
				border-color: #000;
				border-width: 0 0 2px 0;
				border-style: solid;
				height: 46px;
			}
			.el-input /deep/ .el-input__inner:focus {
				border-radius: 0;
				padding: 0 20px;
				color: #666;
				background: #fff;
				flex: 1;
				font-size: 16px;
				border-color: #000;
				border-width: 0 0 2px 0;
				border-style: solid;
				height: 46px;
			}
			.el-input-number {
				width: 100%;
			}
			.el-input-number /deep/ .el-input__inner {
				text-align: center;
				border-radius: 0;
				padding: 0 20px;
				color: #666;
				background: #fff;
				flex: 1;
				font-size: 16px;
				border-color: #000;
				border-width: 0 0 2px 0;
				border-style: solid;
				height: 46px;
			}
			.el-input-number /deep/ .el-input__inner:focus {
				border-radius: 0;
				padding: 0 20px;
				color: #666;
				background: #fff;
				flex: 1;
				font-size: 16px;
				border-color: #000;
				border-width: 0 0 2px 0;
				border-style: solid;
				height: 46px;
			}
			.el-input-number /deep/ .el-input-number__decrease {
				display: none;
			}
			.el-input-number /deep/ .el-input-number__increase {
				display: none;
			}
			.el-select {
				width: 100%;
			}
			.el-select /deep/ .el-input__inner {
				border-radius: 0;
				padding: 0 20px;
				color: #666;
				background: #fff;
				flex: 1;
				font-size: 16px;
				border-color: #000;
				border-width: 0 0 2px 0;
				border-style: solid;
				height: 46px;
			}
			.el-select /deep/ .el-input__inner:focus {
				border-radius: 0;
				padding: 0 20px;
				color: #666;
				background: #fff;
				flex: 1;
				font-size: 16px;
				border-color: #000;
				border-width: 0 0 2px 0;
				border-style: solid;
				height: 46px;
			}
			.el-date-editor {
				width: 100%;
			}
			.el-date-editor /deep/ .el-input__inner {
				border-radius: 0;
				padding: 0 20px 0 30px;
				color: #666;
				background: #fff;
				flex: 1;
				font-size: 16px;
				border-color: #000;
				border-width: 0 0 2px 0;
				border-style: solid;
				height: 46px;
			}
			.el-date-editor /deep/ .el-input__inner:focus {
				border-radius: 0;
				padding: 0 20px 0 30px;
				color: #666;
				background: #fff;
				flex: 1;
				font-size: 16px;
				border-color: #000;
				border-width: 0 0 2px 0;
				border-style: solid;
				height: 46px;
			}
			.el-date-editor.el-input {
				width: 100%;
			}
			/deep/ .el-upload--picture-card {
				background: transparent;
				border: 0;
				border-radius: 0;
				width: auto;
				height: auto;
				line-height: initial;
				vertical-align: middle;
			}
			/deep/ .upload .upload-img {
				border: 1px solid #efeff7;
				cursor: pointer;
				border-radius: 0px;
				margin: 0 45px;
				color: #999;
				background: #fff;
				width: 90px;
				font-size: 24px;
				line-height: 60px;
				text-align: center;
				height: 60px;
			}
			/deep/ .el-upload-list .el-upload-list__item {
				border: 1px solid #efeff7;
				cursor: pointer;
				border-radius: 0px;
				margin: 0 45px;
				color: #999;
				background: #fff;
				width: 90px;
				font-size: 24px;
				line-height: 60px;
				text-align: center;
				height: 60px;
			}
			/deep/ .el-upload .el-icon-plus {
				border: 1px solid #efeff7;
				cursor: pointer;
				border-radius: 0px;
				margin: 0 45px;
				color: #999;
				background: #fff;
				width: 90px;
				font-size: 24px;
				line-height: 60px;
				text-align: center;
				height: 60px;
			}
			/deep/ .el-upload__tip {
				color: #666;
				font-size: 15px;
			}
			/deep/ .el-input__inner::placeholder {
				color: #999;
				font-size: 16px;
			}
			.required {
				position: relative;
			}
			.required::after{
				z-index: 9;
				color: red;
				left: 110px;
				position: inherit;
				content: "*";
				order: -1;
			}
			.editor {
				margin: 0 0 0 10px;
				background: #fff;
				width: 100%;
				height: auto;
			}
			.editor>.avatar-uploader {
				line-height: 0;
				height: 0;
			}
		}
		.list-item.email {
			input {
				border-radius: 0;
				padding: 0 20px;
				color: #666;
				background: #fff;
				flex: 1;
				font-size: 16px;
				border-color: #000;
				border-width: 0 0 2px 0;
				border-style: solid;
				height: 46px;
			}
			input:focus {
				border-radius: 0;
				padding: 0 20px;
				color: #666;
				background: #fff;
				flex: 1;
				font-size: 16px;
				border-color: #000;
				border-width: 0 0 2px 0;
				border-style: solid;
				height: 46px;
			}
			input::placeholder {
				color: #999;
				font-size: 16px;
			}
			button {
				border: 0px solid #efeff7;
				cursor: pointer;
				border-radius: 0 4px 4px 0;
				padding: 0;
				margin: 1px 0 0;
				color: #333;
				background: #0d6efd20;
				width: 150px;
				font-size: 15px;
				height: 46px;
			}
			button:hover {
				opacity: 0.8;
			}
		}
		.register-btn {
			width: 100%;
		}
		.register-btn1 {
			margin: 85px 40px 0  35px;
			display: block;
			width: 100%;
		}
		.register-btn2 {
			width: 100%;
		}
		.r-btn {
			border: 0px solid rgba(0, 0, 0, 1);
			cursor: pointer;
			border-radius: 60px 60px 60px 60px;
			padding: 0 10px;
			margin: 0 0 10px;
			color: #fff;
			background: linear-gradient( 137deg, #57A5FF 0%, #90F4FC 100%);
			font-weight: 700;
			letter-spacing: 2px;
			font-size: 30px;
			min-width: 174px;
			height: 69px;
		}
		.r-btn:hover {
			border: 0px solid rgba(0, 0, 0, 1);
			opacity: 0.8;
		}
		.r-login {
			cursor: pointer;
			padding: 0;
			color: #666;
			display: inline-block;
			text-decoration: underline;
			width: 100%;
			font-size: 15px;
			line-height: 40px;
			text-align: right;
		}
		.r-login:hover {
			opacity: 1;
		}
	}
	.idea-box2 {
		margin: 5px 0 40px;
		background: #fff;
		display: none;
		width: 100%;
		font-size: 16px;
		height: 30px;
		order: -1;
	}
}
	
	::-webkit-scrollbar {
	  display: none;
	}
</style>
