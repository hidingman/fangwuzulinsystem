<template>
	<div class="addEdit-block">
		<el-form
			class="add-update-preview"
			ref="ruleForm"
			:model="ruleForm"
			:rules="rules"
			label-width="180px"
		>
			<template >
				<el-form-item class="input" v-if="type!='info'"  label="房屋名称" prop="fangwumingcheng" >
					<el-input v-model="ruleForm.fangwumingcheng" placeholder="房屋名称" clearable  :readonly="ro.fangwumingcheng"></el-input>
				</el-form-item>
				<el-form-item v-else class="input" label="房屋名称" prop="fangwumingcheng" >
					<el-input v-model="ruleForm.fangwumingcheng" placeholder="房屋名称" readonly></el-input>
				</el-form-item>
				<el-form-item class="upload" v-if="type!='info' && !ro.fangwutupian" label="房屋图片" prop="fangwutupian" >
					<file-upload
						tip="点击上传房屋图片"
						action="file/upload"
						:limit="3"
						:multiple="true"
						:fileUrls="ruleForm.fangwutupian?ruleForm.fangwutupian:''"
						@change="fangwutupianUploadChange"
					></file-upload>
				</el-form-item>
				<el-form-item class="upload" v-else-if="ruleForm.fangwutupian" label="房屋图片" prop="fangwutupian" >
					<img v-if="ruleForm.fangwutupian.substring(0,4)=='http'&&ruleForm.fangwutupian.split(',w').length>1" class="upload-img" style="margin-right:20px;" v-bind:key="index" :src="ruleForm.fangwutupian" width="100" height="100">
					<img v-else-if="ruleForm.fangwutupian.substring(0,4)=='http'" class="upload-img" style="margin-right:20px;" v-bind:key="index" :src="ruleForm.fangwutupian.split(',')[0]" width="100" height="100">
					<img v-else class="upload-img" style="margin-right:20px;" v-bind:key="index" v-for="(item,index) in ruleForm.fangwutupian.split(',')" :src="$base.url+item" width="100" height="100">
				</el-form-item>
				<el-form-item class="input" v-if="type!='info'"  label="房屋户型" prop="fangwuhuxing" >
					<el-input v-model="ruleForm.fangwuhuxing" placeholder="房屋户型" clearable  :readonly="ro.fangwuhuxing"></el-input>
				</el-form-item>
				<el-form-item v-else class="input" label="房屋户型" prop="fangwuhuxing" >
					<el-input v-model="ruleForm.fangwuhuxing" placeholder="房屋户型" readonly></el-input>
				</el-form-item>
				<el-form-item class="date" v-if="type!='info'" label="起始时间" prop="qishishijian" >
					<el-date-picker
						value-format="yyyy-MM-dd HH:mm:ss"
						v-model="ruleForm.qishishijian" 
						type="datetime"
						:readonly="ro.qishishijian"
						placeholder="起始时间"
					></el-date-picker>
				</el-form-item>
				<el-form-item class="input" v-else-if="ruleForm.qishishijian" label="起始时间" prop="qishishijian" >
					<el-input v-model="ruleForm.qishishijian" placeholder="起始时间" readonly></el-input>
				</el-form-item>
				<el-form-item class="date" v-if="type!='info'" label="结束时间" prop="jieshushijian" >
					<el-date-picker
						value-format="yyyy-MM-dd HH:mm:ss"
						v-model="ruleForm.jieshushijian" 
						type="datetime"
						:readonly="ro.jieshushijian"
						placeholder="结束时间"
					></el-date-picker>
				</el-form-item>
				<el-form-item class="input" v-else-if="ruleForm.jieshushijian" label="结束时间" prop="jieshushijian" >
					<el-input v-model="ruleForm.jieshushijian" placeholder="结束时间" readonly></el-input>
				</el-form-item>
				<el-form-item class="input" v-if="type!='info'" label="租赁时间" prop="zulinshijian" >
					<el-input v-model="zulinshijian" placeholder="租赁时间" readonly></el-input>
				</el-form-item>
				<el-form-item class="input" v-else-if="ruleForm.zulinshijian" label="租赁时间" prop="zulinshijian" >
					<el-input v-model="ruleForm.zulinshijian" placeholder="租赁时间" readonly></el-input>
				</el-form-item>
				<el-form-item class="input" v-if="type!='info'"  label="租金价格" prop="zujinjiage" >
					<el-input-number v-model="ruleForm.zujinjiage" placeholder="租金价格" :disabled="ro.zujinjiage"></el-input-number>
				</el-form-item>
				<el-form-item v-else class="input" label="租金价格" prop="zujinjiage" >
					<el-input v-model="ruleForm.zujinjiage" placeholder="租金价格" readonly></el-input>
				</el-form-item>
				<el-form-item class="input" v-if="type!='info'"  label="总计" prop="zongji" >
					<el-input v-model="zongji" placeholder="总计" readonly></el-input>
				</el-form-item>
				<el-form-item class="input" v-else-if="ruleForm.zongji" label="总计" prop="zongji" >
					<el-input v-model="ruleForm.zongji" placeholder="总计" readonly></el-input>
				</el-form-item>
				<el-form-item class="input" v-if="type!='info'"  label="押金" prop="yajin" >
					<el-input-number v-model="ruleForm.yajin" placeholder="押金" :disabled="ro.yajin"></el-input-number>
				</el-form-item>
				<el-form-item v-else class="input" label="押金" prop="yajin" >
					<el-input v-model="ruleForm.yajin" placeholder="押金" readonly></el-input>
				</el-form-item>
				<el-form-item class="input" v-if="type!='info'"  label="积分" prop="jifen" >
					<el-input v-model.number="ruleForm.jifen" placeholder="积分" clearable  :readonly="ro.jifen"></el-input>
				</el-form-item>
				<el-form-item v-else class="input" label="积分" prop="jifen" >
					<el-input v-model="ruleForm.jifen" placeholder="积分" readonly></el-input>
				</el-form-item>
				<el-form-item class="input" v-if="type!='info'"  label="工号" prop="gonghao" >
					<el-input v-model="ruleForm.gonghao" placeholder="工号" clearable  :readonly="ro.gonghao"></el-input>
				</el-form-item>
				<el-form-item v-else class="input" label="工号" prop="gonghao" >
					<el-input v-model="ruleForm.gonghao" placeholder="工号" readonly></el-input>
				</el-form-item>
				<el-form-item class="input" v-if="type!='info'"  label="员工姓名" prop="yuangongxingming" >
					<el-input v-model="ruleForm.yuangongxingming" placeholder="员工姓名" clearable  :readonly="ro.yuangongxingming"></el-input>
				</el-form-item>
				<el-form-item v-else class="input" label="员工姓名" prop="yuangongxingming" >
					<el-input v-model="ruleForm.yuangongxingming" placeholder="员工姓名" readonly></el-input>
				</el-form-item>
				<el-form-item class="input" v-if="type!='info'"  label="账号" prop="zhanghao" >
					<el-input v-model="ruleForm.zhanghao" placeholder="账号" clearable  :readonly="ro.zhanghao"></el-input>
				</el-form-item>
				<el-form-item v-else class="input" label="账号" prop="zhanghao" >
					<el-input v-model="ruleForm.zhanghao" placeholder="账号" readonly></el-input>
				</el-form-item>
				<el-form-item class="input" v-if="type!='info'"  label="姓名" prop="xingming" >
					<el-input v-model="ruleForm.xingming" placeholder="姓名" clearable  :readonly="ro.xingming"></el-input>
				</el-form-item>
				<el-form-item v-else class="input" label="姓名" prop="xingming" >
					<el-input v-model="ruleForm.xingming" placeholder="姓名" readonly></el-input>
				</el-form-item>
			</template>
			<el-form-item class="textarea" v-if="type!='info'" label="租赁备注" prop="zulinbeizhu" >
				<el-input
					style="min-width: 200px; max-width: 600px;"
					type="textarea"
					:rows="8"
					placeholder="租赁备注"
					v-model="ruleForm.zulinbeizhu" >
				</el-input>
			</el-form-item>
			<el-form-item v-else-if="ruleForm.zulinbeizhu" label="租赁备注" prop="zulinbeizhu" >
				<span class="text">{{ruleForm.zulinbeizhu}}</span>
			</el-form-item>
			<el-form-item v-if="type!='info'"  label="租赁详情" prop="zulinxiangqing" >
				<editor 
					style="min-width: 200px; max-width: 600px;"
					v-model="ruleForm.zulinxiangqing" 
					class="editor"
					myQuillEditor="zulinxiangqing"
					action="file/upload">
				</editor>
			</el-form-item>
			<el-form-item v-else-if="ruleForm.zulinxiangqing" label="租赁详情" prop="zulinxiangqing" >
				<span class="text ql-snow ql-editor" v-html="ruleForm.zulinxiangqing"></span>
			</el-form-item>
			<el-form-item class="btn">
				<el-button class="btn3"  v-if="type!='info'" type="success" @click="onSubmit">
					<span class="icon iconfont icon-xihuan"></span>
					提交
				</el-button>
				<el-button class="btn4" v-if="type!='info'" type="success" @click="back()">
					<span class="icon iconfont icon-xihuan"></span>
					取消
				</el-button>
				<el-button class="btn5" v-if="type=='info'" type="success" @click="back()">
					<span class="icon iconfont icon-xihuan"></span>
					返回
				</el-button>
			</el-form-item>
		</el-form>
    

	</div>
</template>
<script>
	import { 
		isNumber,
		isIntNumer,
	} from "@/utils/validate";
	export default {
		data() {
			var validateNumber = (rule, value, callback) => {
				if(!value){
					callback();
				} else if (!isNumber(value)) {
					callback(new Error("请输入数字"));
				} else {
					callback();
				}
			};
			var validateIntNumber = (rule, value, callback) => {
				if(!value){
					callback();
				} else if (!isIntNumer(value)) {
					callback(new Error("请输入整数"));
				} else {
					callback();
				}
			};
			return {
				id: '',
				type: '',
			
			
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
				},

				rules: {
					fangwumingcheng: [
					],
					fangwutupian: [
					],
					fangwuhuxing: [
					],
					qishishijian: [
					],
					jieshushijian: [
					],
					zulinshijian: [
						{ validator: validateIntNumber, trigger: 'blur' },
					],
					zujinjiage: [
						{ validator: validateNumber, trigger: 'blur' },
					],
					zongji: [
						{ validator: validateNumber, trigger: 'blur' },
					],
					yajin: [
						{ validator: validateNumber, trigger: 'blur' },
					],
					jifen: [
						{ validator: validateIntNumber, trigger: 'blur' },
					],
					zulinxiangqing: [
					],
					zulinbeizhu: [
					],
					gonghao: [
					],
					yuangongxingming: [
					],
					zhanghao: [
					],
					xingming: [
					],
					ispay: [
					],
				},
			};
		},
		props: ["parent"],
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



			zongji:{
			
				get: function () {
					return 1*this.ruleForm.zulinshijian*this.ruleForm.zujinjiage+parseFloat(this.ruleForm.yajin==""?0:this.ruleForm.yajin)
				}
			},
		},
		components: {
		},
		created() {
		},
		methods: {
			// 获取日期间隔 单位天
			getFullDay(first, last) {
				let date1 = new Date(first)
				let date2 = new Date(last)
				let a = date1.getTime();
				let b = date2.getTime();
				var count = 0;
				for (var i = a; i <= b; i += 24 * 3600 * 1000) {
					count++;
				}
				return count;
			},
			// 下载
			download(file){
				window.open(`${file}`)
			},
			// 初始化
			init(id,type) {
				if (id) {
					this.id = id;
					this.type = type;
				}
				if(this.type=='info'||this.type=='else'||this.type=='msg'){
					this.info(id);
				}else if(this.type=='logistics'){
					for(let x in this.ro) {
						this.ro[x] = true
					}
					this.logistics=false;
					this.info(id);
				}else if(this.type=='cross'){
					var obj = this.$storage.getObj('crossObj');
					for (var o in obj){
						if(o=='fangwumingcheng'){
							this.ruleForm.fangwumingcheng = obj[o];
							this.ro.fangwumingcheng = true;
							continue;
						}
						if(o=='fangwutupian'){
							this.ruleForm.fangwutupian = obj[o];
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
				// 获取用户信息
				this.$http({
					url: `${this.$storage.get('sessionTable')}/session`,
					method: "get"
				}).then(({ data }) => {
					if (data && data.code === 0) {
						var json = data.data;
						if(((json.zhanghao!=''&&json.zhanghao) || json.zhanghao==0) && this.$storage.get("role")!="管理员"){
							this.ruleForm.zhanghao = json.zhanghao
							this.ro.zhanghao = true;
						}
						if(((json.xingming!=''&&json.xingming) || json.xingming==0) && this.$storage.get("role")!="管理员"){
							this.ruleForm.xingming = json.xingming
							this.ro.xingming = true;
						}
					} else {
						this.$message.error(data.msg);
					}
				});
			
			},
			// 多级联动参数

			info(id) {
				this.$http({
					url: `fangwuzulin/info/${id}`,
					method: "get"
				}).then(({ data }) => {
					if (data && data.code === 0) {
						this.ruleForm = data.data;
						//解决前台上传图片后台不显示的问题
						let reg=new RegExp('../../../upload','g')//g代表全部
						this.ruleForm.zulinxiangqing = this.ruleForm.zulinxiangqing.replace(reg,'../../../fangwuzulinsystem/upload');
					} else {
						this.$message.error(data.msg);
					}
				});
			},

			// 提交
			async onSubmit() {
					if(!this.ruleForm.id) {
						this.ruleForm.ispay = '未支付'
					}
					if(this.ruleForm.zulinshijian==0){
						this.$message.error('租赁时间不能为空')
						return false
					}
					this.ruleForm.zongji = this.zongji
					if(this.ruleForm.fangwutupian!=null) {
						this.ruleForm.fangwutupian = this.ruleForm.fangwutupian.replace(new RegExp(this.$base.url,"g"),"");
					}
					var objcross = this.$storage.getObj('crossObj');
					if(!this.ruleForm.id) {
						delete this.ruleForm.userid
					}
					await this.$refs["ruleForm"].validate(async valid => {
						if (valid) {
							if(this.type=='cross'){
								var statusColumnName = this.$storage.get('statusColumnName');
								var statusColumnValue = this.$storage.get('statusColumnValue');
								if(statusColumnName!='') {
									var obj = this.$storage.getObj('crossObj');
									if(statusColumnName && !statusColumnName.startsWith("[")) {
										for (var o in obj){
											if(o==statusColumnName){
												obj[o] = statusColumnValue;
											}
										}
										var table = this.$storage.get('crossTable');
										await this.$http({
											url: `${table}/update`,
											method: "post",
											data: obj
										}).then(({ data }) => {});
									}
								}
							}
							
							await this.$http({
								url: `fangwuzulin/${!this.ruleForm.id ? "save" : "update"}`,
								method: "post",
								data: this.ruleForm
							}).then(async ({ data }) => {
								if (data && data.code === 0) {
									this.$message({
										message: "操作成功",
										type: "success",
										duration: 1500,
										onClose: () => {
											if(this.isAuth('fangwuzulin','支付')&&this.type=='cross') {
												this.$confirm('是否跳转支付？').then(_ => {
													this.parent.showFlag = true;
													this.parent.addOrUpdateFlag = false;
													this.parent.fangwuzulinCrossAddOrUpdateFlag = false;
													this.$router.push('/fangwuzulin')
												}).catch(_ => {
													this.parent.showFlag = true;
													this.parent.addOrUpdateFlag = false;
													this.parent.fangwuzulinCrossAddOrUpdateFlag = false;
													this.parent.search();
													this.parent.contentStyleChange();
												});
											}else {
												this.parent.showFlag = true;
												this.parent.addOrUpdateFlag = false;
												this.parent.fangwuzulinCrossAddOrUpdateFlag = false;
												this.parent.search();
												this.parent.contentStyleChange();
											}
											
										}
									});
								} else {
									this.$message.error(data.msg);
								}
							});
						}
					});
			},
			// 获取uuid
			getUUID () {
				return new Date().getTime();
			},
			// 返回
			back() {
				this.parent.showFlag = true;
				this.parent.addOrUpdateFlag = false;
				this.parent.fangwuzulinCrossAddOrUpdateFlag = false;
				this.parent.contentStyleChange();
			},
			fangwutupianUploadChange(fileUrls) {
				this.ruleForm.fangwutupian = fileUrls;
			},
		}
	};
</script>
<style lang="scss" scoped>
	.addEdit-block {
		padding: 30px;
	}
	.add-update-preview {
		padding: 40px 80px 80px 0;
		margin: 0 0 0 10px;
		background: #FFFFFF;
		display: flex;
		border-color: #eee;
		border-width: 0px 0 0;
		border-style: solid;
		flex-wrap: wrap;
	}
	.amap-wrapper {
		width: 100%;
		height: 500px;
	}
	
	.search-box {
		position: absolute;
	}
	
	.el-date-editor.el-input {
		width: auto;
	}
	.add-update-preview /deep/ .el-form-item {
		border: 0px solid #eee;
		padding: 0;
		margin: 0 0 26px 0;
		display: inline-block;
		width: 100%;
	}
	.add-update-preview .el-form-item /deep/ .el-form-item__label {
		padding: 0 10px 0 0;
		color: #9E9E9E;
		font-weight: 400;
		width: 180px;
		font-size: 14px;
		line-height: 40px;
		text-align: right;
	}
	
	.add-update-preview .el-form-item /deep/ .el-form-item__content {
		margin-left: 180px;
	}
	.add-update-preview .el-form-item span.text {
		border: 1px solid #E8E8E8;
		padding: 0 10px;
		color: #333;
		word-break: break-all;
		background: none;
		font-weight: 500;
		display: inline-block;
		font-size: 16px;
		min-height: 200px;
		line-height: 40px;
		min-width: 100%;
	}
	
	.add-update-preview .el-input {
		width: 100%;
	}
	.add-update-preview .el-input /deep/ .el-input__inner {
		border: 1px solid #ccc;
		border-radius: 0px;
		padding: 0 12px;
		color: #666;
		width: 100%;
		font-size: 16px;
		min-width: 50%;
		height: 40px;
	}
	.add-update-preview .el-input /deep/ .el-input__inner[readonly="readonly"] {
		border: 1px solid #ccc;
		border-radius: 0px;
		padding: 0 12px;
		color: #666;
		width: 100%;
		font-size: 16px;
		min-width: 50%;
		height: 40px;
	}
	.add-update-preview .el-input-number {
		text-align: left;
		width: 100%;
	}
	.add-update-preview .el-input-number /deep/ .el-input__inner {
		text-align: left;
		border: 1px solid #ccc;
		border-radius: 0px;
		padding: 0 12px;
		color: #666;
		width: 100%;
		font-size: 16px;
		min-width: 50%;
		height: 40px;
	}
	.add-update-preview .el-input-number /deep/ .is-disabled .el-input__inner {
		text-align: left;
		border: 1px solid #ccc;
		border-radius: 0px;
		padding: 0 12px;
		color: #666;
		width: 100%;
		font-size: 16px;
		min-width: 50%;
		height: 40px;
	}
	.add-update-preview .el-input-number /deep/ .el-input-number__decrease {
		display: none;
	}
	.add-update-preview .el-input-number /deep/ .el-input-number__increase {
		display: none;
	}
	.add-update-preview .el-select {
		width: 100%;
	}
	.add-update-preview .el-select /deep/ .el-input__inner {
		border: 1px solid #ccc;
		border-radius: 0px;
		padding: 0 12px;
		color: #666;
		width: 100%;
		font-size: 16px;
		min-width: 50%;
		height: 40px;
	}
	.add-update-preview .el-select /deep/ .is-disabled .el-input__inner {
		border: 1px solid #ccc;
		border-radius: 0px;
		padding: 0 12px;
		color: #666;
		width: 100%;
		font-size: 16px;
		min-width: 50%;
		height: 40px;
	}
	.add-update-preview .el-date-editor {
		width: 100%;
	}
	.add-update-preview .el-date-editor /deep/ .el-input__inner {
		border: 1px solid #ccc;
		border-radius: 0px;
		padding: 0 30px;
		color: #666;
		width: 100%;
		font-size: 16px;
		min-width: 50%;
		height: 40px;
	}
	.add-update-preview .el-date-editor /deep/ .el-input__inner[readonly="readonly"] {
		border: 1px solid #ccc;
		border-radius: 0px;
		padding: 0 30px;
		color: #666;
		width: 100%;
		font-size: 16px;
		min-width: 50%;
		height: 40px;
	}
	.add-update-preview .viewBtn {
		border: 0px solid #ccc;
		cursor: pointer;
		border-radius: 0px;
		padding: 0 15px;
		margin: 0 20px 0 0;
		color: #666;
		background: #fff;
		width: auto;
		font-size: 15px;
		line-height: 34px;
		height: 34px;
		.iconfont {
			margin: 0 2px;
			color: #666;
			font-size: 16px;
			height: 34px;
		}
	}
	.add-update-preview .viewBtn:hover {
		opacity: 0.8;
	}
	.add-update-preview .downBtn {
		border: 0px solid #ccc;
		cursor: pointer;
		border-radius: 0px;
		padding: 0 15px;
		margin: 0 20px 0 0;
		color: #666;
		background: #fff;
		width: auto;
		font-size: 15px;
		line-height: 34px;
		height: 34px;
		.iconfont {
			margin: 0 2px;
			color: #666;
			font-size: 16px;
			height: 34px;
		}
	}
	.add-update-preview .downBtn:hover {
		opacity: 0.8;
	}
	.add-update-preview .unBtn {
		border: 0;
		cursor: not-allowed;
		border-radius: 4px;
		padding: 0 0px;
		margin: 0 20px 0 0;
		outline: none;
		color: #999;
		background: none;
		width: auto;
		font-size: 16px;
		line-height: 40px;
		height: 40px;
		.iconfont {
			margin: 0 2px;
			color: #fff;
			display: none;
			font-size: 14px;
			height: 34px;
		}
	}
	.add-update-preview .unBtn:hover {
		opacity: 0.8;
	}
	.add-update-preview /deep/ .el-upload--picture-card {
		background: transparent;
		border: 0;
		border-radius: 0;
		width: auto;
		height: auto;
		line-height: initial;
		vertical-align: middle;
	}
	
	.add-update-preview /deep/ .upload .upload-img {
		border: 1px solid #ccc;
		cursor: pointer;
		border-radius: 0px;
		color: #666;
		background: #fff;
		width: 90px;
		font-size: 24px;
		line-height: 60px;
		text-align: center;
		height: 60px;
	}
	
	.add-update-preview /deep/ .el-upload-list .el-upload-list__item {
		border: 1px solid #ccc;
		cursor: pointer;
		border-radius: 0px;
		color: #666;
		background: #fff;
		width: 90px;
		font-size: 24px;
		line-height: 60px;
		text-align: center;
		height: 60px;
	}
	
	.add-update-preview /deep/ .el-upload .el-icon-plus {
		border: 1px solid #ccc;
		cursor: pointer;
		border-radius: 0px;
		color: #666;
		background: #fff;
		width: 90px;
		font-size: 24px;
		line-height: 60px;
		text-align: center;
		height: 60px;
	}
	.add-update-preview /deep/ .el-upload__tip {
		color: #666;
		font-size: 15px;
	}
	
	.add-update-preview .el-textarea /deep/ .el-textarea__inner {
		border: 1px solid #ccc;
		border-radius: 0px;
		padding: 0 12px;
		color: #666;
		word-break: break-all;
		width: 100%;
		font-size: 16px;
		min-height: 200px;
		line-height: 24px;
		min-width: 100%;
		height: auto;
	}
	.add-update-preview .el-textarea /deep/ .el-textarea__inner[readonly="readonly"] {
				border: 1px solid #ccc;
				border-radius: 0px;
				padding: 0 12px;
				color: #666;
				word-break: break-all;
				width: 100%;
				font-size: 16px;
				min-height: 200px;
				line-height: 24px;
				min-width: 100%;
				height: auto;
			}
	.add-update-preview .el-form-item.btn {
		padding: 0;
		margin: 20px 0 0;
		.btn1 {
			border: 0px solid #ccc;
			cursor: pointer;
			border-radius: 4px;
			padding: 0 10px;
			margin: 0 10px 0 0;
			color: #fff;
			background: #5BAAFF;
			width: auto;
			font-size: 16px;
			min-width: 110px;
			height: 40px;
			.iconfont {
				margin: 0 2px;
				color: #fff;
				display: none;
				font-size: 14px;
				height: 40px;
			}
		}
		.btn1:hover {
			opacity: 0.8;
		}
		.btn2 {
			border: 0px solid #ccc;
			cursor: pointer;
			border-radius: 4px;
			padding: 0 10px;
			margin: 0 10px 0 0;
			color: #fff;
			background: #60DFE4;
			width: auto;
			font-size: 16px;
			min-width: 110px;
			height: 40px;
			.iconfont {
				margin: 0 2px;
				color: #fff;
				display: none;
				font-size: 14px;
				height: 34px;
			}
		}
		.btn2:hover {
			opacity: 0.8;
		}
		.btn3 {
			border: 0px solid #ccc;
			cursor: pointer;
			border-radius: 4px;
			padding: 0 10px;
			margin: 0 10px 0 0;
			color: #fff;
			background: #60E495;
			width: auto;
			font-size: 16px;
			min-width: 110px;
			height: 40px;
			.iconfont {
				margin: 0 2px;
				color: #fff;
				display: none;
				font-size: 14px;
				height: 40px;
			}
		}
		.btn3:hover {
			opacity: 0.8;
		}
		.btn4 {
			border: 0px solid #ccc;
			cursor: pointer;
			border-radius: 4px;
			padding: 0 10px;
			margin: 0 10px 0 0;
			color: #fff;
			background: #C3E460;
			width: auto;
			font-size: 16px;
			min-width: 110px;
			height: 40px;
			.iconfont {
				margin: 0 2px;
				color: #fff;
				display: none;
				font-size: 14px;
				height: 40px;
			}
		}
		.btn4:hover {
			opacity: 0.8;
		}
		.btn5 {
			border: 0px solid #ccc;
			cursor: pointer;
			border-radius: 4px;
			padding: 0 10px;
			margin: 0 10px 0 0;
			color: #fff;
			background: #E4B860;
			width: auto;
			font-size: 16px;
			min-width: 110px;
			height: 40px;
			.iconfont {
				margin: 0 2px;
				color: #fff;
				display: none;
				font-size: 14px;
				height: 40px;
			}
		}
		.btn5:hover {
			opacity: 0.8;
		}
	}
</style>
