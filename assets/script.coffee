jQuery ->
    # bootstrapping
    $.get "/assets/template.html", (template) ->
        document.write template
        initTemplates()
        main()

    answerRowTemplate = undefined

    initTemplates = ->
        answerRowTemplate = $("#answer-row").clone()
        $("#answer-table tbody *").remove()

    main = ->
        showScreen("screen-index")
        renderAnswers()
        setUpFilter()
        linkDetail()

    showScreen = (id)->
        $("#screens > div").hide()
        $("#" + id).show()    

    renderAnswers = ->
        $.get "/answers", (answers) ->
            $(answers).each (index, answer) ->
                answerRow = answerRowTemplate.clone()
                answerRow
                    .find(".No").text(answer.No).end()
                    .find(".ChallengerNo").text(answer.ChallengerNo).end()
                    .find(".ChallengeCount").text(answer.ChallengeCount).end()
                    .find(".Nickname").text(answer.Nickname).end()
                    .find(".Gender").text(answer.Gender).end()
                    .find(".Age").text(answer.Age).end()
                    .find(".ChallengedOn").text(answer.ChallengedOn).end()
                    .find(".Feedbacked").text(answer.Feedbacked).end()
                    .find("a.detail").attr("data-no", answer.No).end()
                    .attr("data-feedbacked", answer.Feedbacked)
                    .appendTo("#answer-table tbody")

    renderAnswer = (number)->
        $.get("/answers/#{number}")
            .success (answer)->
                showScreen("screen-detail")
                $("#screen-detail")
                    .find(".No").text(answer.No).end()
                    .find(".ChallengerNo").text(answer.ChallengerNo).end()
                    .find(".ChallengeCount").text(answer.ChallengeCount).end()
                    .find(".Nickname").text(answer.Nickname).end()
                    .find(".Gender").text(answer.Gender).end()
                    .find(".Age").text(answer.Age).end()
                    .find(".ChallengedOn").text(answer.ChallengedOn).end()
                    .find(".Feedbacked").text(answer.Feedbacked).end()
                    .find(".AnswerText").text(answer.AnswerText).end()

            .error ()->
                alert("failed to load data")

    setUpFilter = ->
        $("#filter-by-feedback a").on "click", ->
            filterBy = $(this).data("filter")

            $("#filter-by-feedback li").removeClass("active")
            $(this).parent().addClass("active")

            if filterBy is "done"
                $("#answer-table tbody tr:not([data-feedbacked=済])").hide()
                $("#answer-table tbody tr[data-feedbacked=済]").show()
                return false
            if filterBy is "yet"
                $("#answer-table tbody tr:not([data-feedbacked=未])").hide()
                $("#answer-table tbody tr[data-feedbacked=未]").show()
                return false    
    
            $("#answer-table tbody tr").show()
            return false

    linkDetail = ->
        $(document).on "click", "a.detail", ->
            renderAnswer($(this).data("no"))
            return false
