'use strict';

/* Controllers */

angular.module('myApp.controllers', []);

angular.module('myApp').
    controller('IndexesCtrl', IndexesCtrl).
    controller('IndexCtrl', IndexCtrl).
    controller('IndexNewCtrl', IndexNewCtrl).
    controller('MonitorCtrl', MonitorCtrl).
    controller('SearchCtrl', SearchCtrl);

angular.module('myApp').
    controller('BleveAnalyzerModalCtrl', BleveAnalyzerModalCtrl).
    controller('BleveCharFilterModalCtrl', BleveCharFilterModalCtrl).
    controller('BleveTokenizerModalCtrl', BleveTokenizerModalCtrl).
    controller('BleveTokenFilterModalCtrl', BleveTokenFilterModalCtrl).
    controller('BleveWordListModalCtrl', BleveWordListModalCtrl);
