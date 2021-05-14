"use strict";

const Generator = require("yeoman-generator");
const chalk = require("chalk");
const yosay = require("yosay");

module.exports = class extends Generator {
  async prompting() {
    // Have Yeoman greet the user.
    this.log(
      yosay(
        `Welcome to the fabulous ${chalk.red(
          "generator-my-go-cgi-project"
        )} generator!`
      )
    );

    this.props = [];

    this.props.push(
      await this.prompt([
        {
          type: "input",
          name: "project_name",
          message: "Input project name.",
          default: "my_go_cgi_project"
        }
      ])
    );

    this.props = Object.assign({}, ...this.props);
  }

  writing() {
    console.log(`pkg/${this.props.project_name}`);

    this._copyTarget([
      ["_gitignore", ".gitignore", null],
      ["go.mod", "go.mod", this.props],
      ["readme.md", "readme.md", this.props],

      ["pkg/__projectname__", `pkg/${this.props.project_name}`, this.props],
      ["cmd/__projectname__", `cmd/${this.props.project_name}`, this.props],
      ["cmd/garg", "cmd/garg", null]
    ]);
  }

  _copyTarget(targets) {
    for (const t of targets) {
      if (t[2]) {
        this.fs.copyTpl(
          this.templatePath(t[0]),
          this.destinationPath(t[1]),
          t[2]
        );
      } else {
        this.fs.copy(this.templatePath(t[0]), this.destinationPath(t[1]));
      }
    }
  }

  install() {
    this.spawnCommand("go", ["mod", "download"]);
  }
};
