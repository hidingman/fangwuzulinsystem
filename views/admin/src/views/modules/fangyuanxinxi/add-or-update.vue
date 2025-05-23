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
				<el-form-item class="select" v-if="type!='info'"  label="房屋户型" prop="fangwuhuxing" >
					<el-select :disabled="ro.fangwuhuxing" v-model="ruleForm.fangwuhuxing" placeholder="请选择房屋户型" >
						<el-option
							v-for="(item,index) in fangwuhuxingOptions"
							v-bind:key="index"
							:label="item"
							:value="item">
						</el-option>
					</el-select>
				</el-form-item>
				<el-form-item v-else class="input" label="房屋户型" prop="fangwuhuxing" >
					<el-input v-model="ruleForm.fangwuhuxing"
						placeholder="房屋户型" readonly></el-input>
				</el-form-item>
				<el-form-item class="input" v-if="type!='info'"  label="房屋地址" prop="fangwudizhi" >
					<el-input v-model="ruleForm.fangwudizhi" placeholder="房屋地址" clearable  :readonly="ro.fangwudizhi"></el-input>
				</el-form-item>
				<el-form-item v-else class="input" label="房屋地址" prop="fangwudizhi" >
					<el-input v-model="ruleForm.fangwudizhi" placeholder="房屋地址" readonly></el-input>
				</el-form-item>
				<el-form-item class="input" v-if="type!='info'"  label="面积" prop="mianji" >
					<el-input v-model="ruleForm.mianji" placeholder="面积" clearable  :readonly="ro.mianji"></el-input>
				</el-form-item>
				<el-form-item v-else class="input" label="面积" prop="mianji" >
					<el-input v-model="ruleForm.mianji" placeholder="面积" readonly></el-input>
				</el-form-item>
				<el-form-item class="input" v-if="type!='info'"  label="租金价格" prop="zujinjiage" >
					<el-input-number v-model="ruleForm.zujinjiage" placeholder="租金价格" :disabled="ro.zujinjiage"></el-input-number>
				</el-form-item>
				<el-form-item v-else class="input" label="租金价格" prop="zujinjiage" >
					<el-input v-model="ruleForm.zujinjiage" placeholder="租金价格" readonly></el-input>
				</el-form-item>
				<el-form-item class="select" v-if="type!='info'"  label="押金方式" prop="yajinfangshi" >
					<el-select :disabled="ro.yajinfangshi" v-model="ruleForm.yajinfangshi" placeholder="请选择押金方式" >
						<el-option
							v-for="(item,index) in yajinfangshiOptions"
							v-bind:key="index"
							:label="item"
							:value="item">
						</el-option>
					</el-select>
				</el-form-item>
				<el-form-item v-else class="input" label="押金方式" prop="yajinfangshi" >
					<el-input v-model="ruleForm.yajinfangshi"
						placeholder="押金方式" readonly></el-input>
				</el-form-item>
				<el-form-item class="select" v-if="type!='info'"  label="付款方式" prop="fukuanfangshi" >
					<el-select :disabled="ro.fukuanfangshi" v-model="ruleForm.fukuanfangshi" placeholder="请选择付款方式" >
						<el-option
							v-for="(item,index) in fukuanfangshiOptions"
							v-bind:key="index"
							:label="item"
							:value="item">
						</el-option>
					</el-select>
				</el-form-item>
				<el-form-item v-else class="input" label="付款方式" prop="fukuanfangshi" >
					<el-input v-model="ruleForm.fukuanfangshi"
						placeholder="付款方式" readonly></el-input>
				</el-form-item>
				<el-form-item class="input" v-if="type!='info'"  label="房屋朝向" prop="fangwuchaoxiang" >
					<el-input v-model="ruleForm.fangwuchaoxiang" placeholder="房屋朝向" clearable  :readonly="ro.fangwuchaoxiang"></el-input>
				</el-form-item>
				<el-form-item v-else class="input" label="房屋朝向" prop="fangwuchaoxiang" >
					<el-input v-model="ruleForm.fangwuchaoxiang" placeholder="房屋朝向" readonly></el-input>
				</el-form-item>
				<el-form-item class="input" v-if="type!='info'"  label="押金" prop="yajin" >
					<el-input-number v-model="ruleForm.yajin" placeholder="押金" :disabled="ro.yajin"></el-input-number>
				</el-form-item>
				<el-form-item v-else class="input" label="押金" prop="yajin" >
					<el-input v-model="ruleForm.yajin" placeholder="押金" readonly></el-input>
				</el-form-item>
				<el-form-item class="input" v-if="type!='info'"  label="小区" prop="xiaoqu" >
					<el-input v-model="ruleForm.xiaoqu" placeholder="小区" clearable  :readonly="ro.xiaoqu"></el-input>
				</el-form-item>
				<el-form-item v-else class="input" label="小区" prop="xiaoqu" >
					<el-input v-model="ruleForm.xiaoqu" placeholder="小区" readonly></el-input>
				</el-form-item>
				<el-form-item class="upload" v-if="type!='info'&& !ro.fangwushipin" label="房屋视频" prop="fangwushipin" >
					<file-upload
						tip="点击上传房屋视频"
						action="file/upload"
						:limit="1"
						:type="2"
						:multiple="true"
						:fileUrls="ruleForm.fangwushipin?ruleForm.fangwushipin:''"
						@change="fangwushipinUploadChange"
					></file-upload>
				</el-form-item> 
				<el-form-item v-else-if="ruleForm.fangwushipin" label="房屋视频" prop="fangwushipin" >
					<el-button class="viewBtn" type="text" size="small" @click="download($base.url+ruleForm.fangwushipin)">
						<span class="icon iconfont icon-chakan14"></span>
						预览
					</el-button>
				</el-form-item>
				<el-form-item v-else-if="!ruleForm.fangwushipin" label="房屋视频" prop="fangwushipin" >
					<el-button class="unBtn" type="text" size="small">
						<span class="icon iconfont icon-xihuan"></span>
						暂无
					</el-button>
				</el-form-item>
				<el-form-item class="input" v-if="type!='info'"  label="楼栋单元" prop="loudongdanyuan" >
					<el-input v-model="ruleForm.loudongdanyuan" placeholder="楼栋单元" clearable  :readonly="ro.loudongdanyuan"></el-input>
				</el-form-item>
				<el-form-item v-else class="input" label="楼栋单元" prop="loudongdanyuan" >
					<el-input v-model="ruleForm.loudongdanyuan" placeholder="楼栋单元" readonly></el-input>
				</el-form-item>
				<el-form-item class="input" v-if="type!='info'"  label="房号" prop="fanghao" >
					<el-input v-model="ruleForm.fanghao" placeholder="房号" clearable  :readonly="ro.fanghao"></el-input>
				</el-form-item>
				<el-form-item v-else class="input" label="房号" prop="fanghao" >
					<el-input v-model="ruleForm.fanghao" placeholder="房号" readonly></el-input>
				</el-form-item>
				<el-form-item class="input" v-if="type!='info'"  label="积分" prop="jifen" >
					<el-input v-model.number="ruleForm.jifen" placeholder="积分" clearable  :readonly="ro.jifen"></el-input>
				</el-form-item>
				<el-form-item v-else class="input" label="积分" prop="jifen" >
					<el-input v-model="ruleForm.jifen" placeholder="积分" readonly></el-input>
				</el-form-item>
				<el-form-item class="input" v-if="type!='info'"  label="房屋结构" prop="fangwujiegou" >
					<el-input v-model="ruleForm.fangwujiegou" placeholder="房屋结构" clearable  :readonly="ro.fangwujiegou"></el-input>
				</el-form-item>
				<el-form-item v-else class="input" label="房屋结构" prop="fangwujiegou" >
					<el-input v-model="ruleForm.fangwujiegou" placeholder="房屋结构" readonly></el-input>
				</el-form-item>
				<el-form-item class="select" v-if="type!='info'"  label="房屋状态" prop="fangwuzhuangtai" >
					<el-select :disabled="ro.fangwuzhuangtai" v-model="ruleForm.fangwuzhuangtai" placeholder="请选择房屋状态" >
						<el-option
							v-for="(item,index) in fangwuzhuangtaiOptions"
							v-bind:key="index"
							:label="item"
							:value="item">
						</el-option>
					</el-select>
				</el-form-item>
				<el-form-item v-else class="input" label="房屋状态" prop="fangwuzhuangtai" >
					<el-input v-model="ruleForm.fangwuzhuangtai"
						placeholder="房屋状态" readonly></el-input>
				</el-form-item>
				<el-form-item class="input" v-if="type!='info'"  label="房产证编号" prop="fangchanzhengbianhao" >
					<el-input v-model="ruleForm.fangchanzhengbianhao" placeholder="房产证编号" clearable  :readonly="ro.fangchanzhengbianhao"></el-input>
				</el-form-item>
				<el-form-item v-else class="input" label="房产证编号" prop="fangchanzhengbianhao" >
					<el-input v-model="ruleForm.fangchanzhengbianhao" placeholder="房产证编号" readonly></el-input>
				</el-form-item>
				<el-form-item class="upload" v-if="type!='info' && !ro.fangchanzhengzhaopian" label="房产证照片" prop="fangchanzhengzhaopian" >
					<file-upload
						tip="点击上传房产证照片"
						action="file/upload"
						:limit="3"
						:multiple="true"
						:fileUrls="ruleForm.fangchanzhengzhaopian?ruleForm.fangchanzhengzhaopian:''"
						@change="fangchanzhengzhaopianUploadChange"
					></file-upload>
				</el-form-item>
				<el-form-item class="upload" v-else-if="ruleForm.fangchanzhengzhaopian" label="房产证照片" prop="fangchanzhengzhaopian" >
					<img v-if="ruleForm.fangchanzhengzhaopian.substring(0,4)=='http'&&ruleForm.fangchanzhengzhaopian.split(',w').length>1" class="upload-img" style="margin-right:20px;" v-bind:key="index" :src="ruleForm.fangchanzhengzhaopian" width="100" height="100">
					<img v-else-if="ruleForm.fangchanzhengzhaopian.substring(0,4)=='http'" class="upload-img" style="margin-right:20px;" v-bind:key="index" :src="ruleForm.fangchanzhengzhaopian.split(',')[0]" width="100" height="100">
					<img v-else class="upload-img" style="margin-right:20px;" v-bind:key="index" v-for="(item,index) in ruleForm.fangchanzhengzhaopian.split(',')" :src="$base.url+item" width="100" height="100">
				</el-form-item>
				<el-form-item class="input" v-if="type!='info'"  label="房主姓名" prop="fangzhuxingming" >
					<el-input v-model="ruleForm.fangzhuxingming" placeholder="房主姓名" clearable  :readonly="ro.fangzhuxingming"></el-input>
				</el-form-item>
				<el-form-item v-else class="input" label="房主姓名" prop="fangzhuxingming" >
					<el-input v-model="ruleForm.fangzhuxingming" placeholder="房主姓名" readonly></el-input>
				</el-form-item>
				<el-form-item class="input" v-if="type!='info'"  label="房主身份证" prop="fangzhushenfenzheng" >
					<el-input v-model="ruleForm.fangzhushenfenzheng" placeholder="房主身份证" clearable  :readonly="ro.fangzhushenfenzheng"></el-input>
				</el-form-item>
				<el-form-item v-else class="input" label="房主身份证" prop="fangzhushenfenzheng" >
					<el-input v-model="ruleForm.fangzhushenfenzheng" placeholder="房主身份证" readonly></el-input>
				</el-form-item>
				<el-form-item class="input" v-if="type!='info'"  label="房主电话" prop="fangzhudianhua" >
					<el-input v-model="ruleForm.fangzhudianhua" placeholder="房主电话" clearable  :readonly="ro.fangzhudianhua"></el-input>
				</el-form-item>
				<el-form-item v-else class="input" label="房主电话" prop="fangzhudianhua" >
					<el-input v-model="ruleForm.fangzhudianhua" placeholder="房主电话" readonly></el-input>
				</el-form-item>
				<el-form-item class="date" v-if="type!='info'" label="发布时间" prop="fabushijian" >
					<el-date-picker
						value-format="yyyy-MM-dd HH:mm:ss"
						v-model="ruleForm.fabushijian" 
						type="datetime"
						:readonly="ro.fabushijian"
						placeholder="发布时间"
					></el-date-picker>
				</el-form-item>
				<el-form-item class="input" v-else-if="ruleForm.fabushijian" label="发布时间" prop="fabushijian" >
					<el-input v-model="ruleForm.fabushijian" placeholder="发布时间" readonly></el-input>
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
			</template>
			<el-form-item v-if="type!='info'"  label="房屋详情" prop="fangwuxiangqing" >
				<editor 
					style="min-width: 200px; max-width: 600px;"
					v-model="ruleForm.fangwuxiangqing" 
					class="editor"
					myQuillEditor="fangwuxiangqing"
					action="file/upload">
				</editor>
			</el-form-item>
			<el-form-item v-else-if="ruleForm.fangwuxiangqing" label="房屋详情" prop="fangwuxiangqing" >
				<span class="text ql-snow ql-editor" v-html="ruleForm.fangwuxiangqing"></span>
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
		isMobile,
		checkIdCard,
	} from "@/utils/validate";
	export default {
		data() {
			var validateIdCard = (rule, value, callback) => {
				if(!value){
					callback();
				} else if (!checkIdCard(value)) {
					callback(new Error("请输入正确的身份证号码"));
				} else {
					callback();
				}
			};
			var validateMobile = (rule, value, callback) => {
				if(!value){
					callback();
				} else if (!isMobile(value)) {
					callback(new Error("请输入正确的手机号码"));
				} else {
					callback();
				}
			};
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
					jifen: Number('5'),
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
					clicktime: '',
				},
				fangwuhuxingOptions: [],
				yajinfangshiOptions: [],
				fukuanfangshiOptions: [],
				fangwuzhuangtaiOptions: [],

				rules: {
					fangwumingcheng: [
					],
					fangwutupian: [
					],
					fangwuhuxing: [
					],
					fangwudizhi: [
					],
					mianji: [
					],
					zujinjiage: [
						{ validator: validateNumber, trigger: 'blur' },
					],
					yajinfangshi: [
					],
					fukuanfangshi: [
					],
					fangwuchaoxiang: [
					],
					yajin: [
						{ validator: validateNumber, trigger: 'blur' },
					],
					xiaoqu: [
					],
					fangwushipin: [
					],
					loudongdanyuan: [
					],
					fanghao: [
					],
					jifen: [
						{ validator: validateIntNumber, trigger: 'blur' },
					],
					fangwujiegou: [
					],
					fangwuzhuangtai: [
					],
					fangchanzhengbianhao: [
					],
					fangchanzhengzhaopian: [
					],
					fangzhuxingming: [
					],
					fangzhushenfenzheng: [
						{ validator: validateIdCard, trigger: 'blur' },
					],
					fangzhudianhua: [
						{ validator: validateMobile, trigger: 'blur' },
					],
					fangwuxiangqing: [
					],
					fabushijian: [
					],
					gonghao: [
					],
					yuangongxingming: [
					],
					thumbsupnum: [
						{ validator: validateIntNumber, trigger: 'blur' },
					],
					crazilynum: [
						{ validator: validateIntNumber, trigger: 'blur' },
					],
					clicktime: [
					],
					clicknum: [
						{ validator: validateIntNumber, trigger: 'blur' },
					],
					discussnum: [
						{ validator: validateIntNumber, trigger: 'blur' },
					],
					storeupnum: [
						{ validator: validateIntNumber, trigger: 'blur' },
					],
				},
			};
		},
		props: ["parent"],
		computed: {



		},
		components: {
		},
		created() {
		},
		methods: {
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
							this.ruleForm.fangchanzhengzhaopian = obj[o];
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
					this.ruleForm.jifen = Number('5'); 				}
				// 获取用户信息
				this.$http({
					url: `${this.$storage.get('sessionTable')}/session`,
					method: "get"
				}).then(({ data }) => {
					if (data && data.code === 0) {
						var json = data.data;
						if(this.$storage.get("role")!="管理员") {
							this.ro.jifen = true;
						}
						if(((json.gonghao!=''&&json.gonghao) || json.gonghao==0) && this.$storage.get("role")!="管理员"){
							this.ruleForm.gonghao = json.gonghao
							this.ro.gonghao = true;
						}
						if(((json.yuangongxingming!=''&&json.yuangongxingming) || json.yuangongxingming==0) && this.$storage.get("role")!="管理员"){
							this.ruleForm.yuangongxingming = json.yuangongxingming
							this.ro.yuangongxingming = true;
						}
					} else {
						this.$message.error(data.msg);
					}
				});
				this.$http({
					url: `option/fangwuhuxing/fangwuhuxing`,
					method: "get"
				}).then(({ data }) => {
					if (data && data.code === 0) {
						this.fangwuhuxingOptions = data.data;
					} else {
						this.$message.error(data.msg);
					}
				});
				this.yajinfangshiOptions = "押一付一,押一付三,半年付,年付".split(',')
				this.fukuanfangshiOptions = "转账,现金".split(',')
				this.fangwuzhuangtaiOptions = "未租,已租".split(',')
			
			},
			// 多级联动参数

			info(id) {
				this.$http({
					url: `fangyuanxinxi/info/${id}`,
					method: "get"
				}).then(({ data }) => {
					if (data && data.code === 0) {
						this.ruleForm = data.data;
						//解决前台上传图片后台不显示的问题
						let reg=new RegExp('../../../upload','g')//g代表全部
						this.ruleForm.fangwuxiangqing = this.ruleForm.fangwuxiangqing.replace(reg,'../../../fangwuzulinsystem/upload');
					} else {
						this.$message.error(data.msg);
					}
				});
			},

			// 提交
			async onSubmit() {
					if(this.ruleForm.fangwutupian!=null) {
						this.ruleForm.fangwutupian = this.ruleForm.fangwutupian.replace(new RegExp(this.$base.url,"g"),"");
					}
					if(this.ruleForm.fangwushipin!=null) {
						this.ruleForm.fangwushipin = this.ruleForm.fangwushipin.replace(new RegExp(this.$base.url,"g"),"");
					}
					if(this.ruleForm.fangchanzhengzhaopian!=null) {
						this.ruleForm.fangchanzhengzhaopian = this.ruleForm.fangchanzhengzhaopian.replace(new RegExp(this.$base.url,"g"),"");
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
								url: `fangyuanxinxi/${!this.ruleForm.id ? "save" : "update"}`,
								method: "post",
								data: this.ruleForm
							}).then(async ({ data }) => {
								if (data && data.code === 0) {
									this.$message({
										message: "操作成功",
										type: "success",
										duration: 1500,
										onClose: () => {
											this.parent.showFlag = true;
											this.parent.addOrUpdateFlag = false;
											this.parent.fangyuanxinxiCrossAddOrUpdateFlag = false;
											this.parent.search();
											this.parent.contentStyleChange();
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
				this.parent.fangyuanxinxiCrossAddOrUpdateFlag = false;
				this.parent.contentStyleChange();
			},
			fangwutupianUploadChange(fileUrls) {
				this.ruleForm.fangwutupian = fileUrls;
			},
			fangwushipinUploadChange(fileUrls) {
				this.ruleForm.fangwushipin = fileUrls;
			},
			fangchanzhengzhaopianUploadChange(fileUrls) {
				this.ruleForm.fangchanzhengzhaopian = fileUrls;
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
