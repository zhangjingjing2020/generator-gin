import Generator from 'yeoman-generator';
import path from 'path';
import fs from 'fs';

export default class extends Generator {
    async prompting() {
        this.answers = await this.prompt([{
                type: 'input',
                name: 'projectName',
                message: '请输入项目名称：',
                default: this.appname,
            },
            {
                type: 'input',
                name: 'moduleName',
                message: '请输入 Go module 名称：',
                default: (answers) => `${answers.projectName}`,
            },
            {
                type: 'input',
                name: 'outputPath',
                message: '请输入生成目录（默认当前目录）：',
                default: process.cwd(),
            },
        ]);
    }

    writing() {
        const targetDir = path.join(this.answers.outputPath, this.answers.projectName);

        // 确保目录存在
        fs.mkdirSync(targetDir, { recursive: true });

        // 1渲染 go.mod.tpl → go.mod
        this.fs.copyTpl(
            this.templatePath('go.mod.tpl'),
            path.join(targetDir, 'go.mod'),
            this.answers
        );

        // 2复制其他模板文件（**/* 匹配所有文件，包括隐藏文件）
        this.fs.copyTpl(
            this.templatePath('**/*'),
            targetDir,
            this.answers, {}, { globOptions: { dot: true, ignore: ['go.mod.tpl'] } } // 支持 .gitignore 等隐藏文件
        );

        // 3删除模板文件
        const filesToDelete = ['go.mod.tpl'];
        filesToDelete.forEach(file => {
            const filePath = path.join(targetDir, file);
            if (fs.existsSync(filePath)) {
                fs.unlinkSync(filePath);
            }
        });
    }

    end() {
        this.log(`✅ 项目 ${this.answers.projectName} 已生成到 ${this.answers.outputPath}`);
    }
}