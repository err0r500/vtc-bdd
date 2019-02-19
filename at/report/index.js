var reporter = require('cucumber-html-reporter');

var options = {
    theme: 'simple',
    jsonFile: 'report.json',
    output: 'report.html',
    reportSuiteAsScenarios: true,
    launchReport: false,
};

reporter.generate(options);

