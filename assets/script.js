// Generated by CoffeeScript 1.6.3
(function() {
  jQuery(function() {
    var answerRowTemplate, initTemplates, linkDetail, main, renderAnswer, renderAnswers, setUpFilter, showScreen;
    $.get("/assets/template.html", function(template) {
      document.write(template);
      initTemplates();
      return main();
    });
    answerRowTemplate = void 0;
    initTemplates = function() {
      answerRowTemplate = $("#answer-row").clone();
      return $("#answer-table tbody *").remove();
    };
    main = function() {
      showScreen("screen-index");
      renderAnswers();
      setUpFilter();
      return linkDetail();
    };
    showScreen = function(id) {
      $("#screens > div").hide();
      return $("#" + id).show();
    };
    renderAnswers = function() {
      return $.get("/answers", function(answers) {
        return $(answers).each(function(index, answer) {
          var answerRow;
          answerRow = answerRowTemplate.clone();
          return answerRow.find(".No").text(answer.No).end().find(".ChallengerNo").text(answer.ChallengerNo).end().find(".ChallengeCount").text(answer.ChallengeCount).end().find(".Nickname").text(answer.Nickname).end().find(".Gender").text(answer.Gender).end().find(".Age").text(answer.Age).end().find(".ChallengedOn").text(answer.ChallengedOn).end().find(".Feedbacked").text(answer.Feedbacked).end().find("a.detail").attr("data-no", answer.No).end().attr("data-feedbacked", answer.Feedbacked).appendTo("#answer-table tbody");
        });
      });
    };
    renderAnswer = function(number) {
      return $.get("/answers/" + number).success(function(answer) {
        showScreen("screen-detail");
        return $("#screen-detail").find(".No").text(answer.No).end().find(".ChallengerNo").text(answer.ChallengerNo).end().find(".ChallengeCount").text(answer.ChallengeCount).end().find(".Nickname").text(answer.Nickname).end().find(".Gender").text(answer.Gender).end().find(".Age").text(answer.Age).end().find(".ChallengedOn").text(answer.ChallengedOn).end().find(".Feedbacked").text(answer.Feedbacked).end().find(".AnswerText").text(answer.AnswerText).end();
      }).error(function() {
        return alert("failed to load data");
      });
    };
    setUpFilter = function() {
      return $("#filter-by-feedback a").on("click", function() {
        var filterBy;
        filterBy = $(this).data("filter");
        $("#filter-by-feedback li").removeClass("active");
        $(this).parent().addClass("active");
        if (filterBy === "done") {
          $("#answer-table tbody tr:not([data-feedbacked=済])").hide();
          $("#answer-table tbody tr[data-feedbacked=済]").show();
          return false;
        }
        if (filterBy === "yet") {
          $("#answer-table tbody tr:not([data-feedbacked=未])").hide();
          $("#answer-table tbody tr[data-feedbacked=未]").show();
          return false;
        }
        $("#answer-table tbody tr").show();
        return false;
      });
    };
    return linkDetail = function() {
      return $(document).on("click", "a.detail", function() {
        renderAnswer($(this).data("no"));
        return false;
      });
    };
  });

}).call(this);