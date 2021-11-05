<script setup>

</script>

<template>
  <div class="topictitle">
    <h1 class="title">Documenting a Swift Framework or Package</h1>
  </div>

  <div class="container content-grid">
    <div class="description">
      <div class="abstract">Create rich and engaging documentation from your in-source comments, and add a documentation
        catalog to your project to provide additional content.
      </div>
    </div>
    <div class="summary">On This Page</div>
    <div class="primary-content">
      <h2>Overview</h2>

      <p>DocC, or Documentation Compiler, makes it easy for you to produce documentation for your Swift frameworks and
        packages. The compiler builds your documentation by combining comments you write in-source with extension files,
        articles, and other resources that live alongside your project in Xcode, allowing you to create rich and
        engaging documentation for developers.</p>

      <p>With DocC, you provide a combination of reference and conceptual content, and connect it together using
        powerful organization and linking capabilities. The compiler integrates directly with Xcode to enhance your
        existing workflow, including code completion, Quick Help, and more. And because you write documentation directly
        in source, you can use the tools you’re already familiar with, such as Git, to track changes you make.</p>

      <h3
          id="Incorporate-Documentation-into-Your-Build-Process"
      >Incorporate Documentation into Your Build Process</h3>

      <ol>
        <li>
          <p>In Xcode, select your framework’s project in the Project navigator.</p>
        </li>
        <li>
          <p>Select the framework’s target in the project editor.</p>
        </li>
        <li>
          <p>Click the Build Settings tab.</p>
        </li>
        <li>
          <p>Enter “build documentation” in the search box to locate the Build Documentation during ‘Build’ setting.</p>
        </li>
        <li>
          <p>Choose Yes from the setting’s pop-up button to enable the build setting.</p>
        </li>
      </ol>

      <img
          src="https://docs-assets.developer.apple.com/published/fdd6b5d865669db0f976405dc24b10ac/11800/doc-viewer-parameters@2x.png">


      <div class="code-listing">
        <pre>
123123
        </pre>

      </div>

      <aside aria-label="note" class="note">
        <p class="label">Note</p>
        <p>For existing framework projects, the build setting appears only after you add a documentation catalog. You
          can’t use the build setting with a Swift package.</p>
      </aside>

      <aside aria-label="important" class="important">
        <p class="label">Important</p>
        <p>For existing framework projects, the build setting appears only after you add a documentation catalog. You
          can’t use the build setting with a Swift package.</p>
      </aside>

      <aside aria-label="deprecated" class="deprecated">
        <p class="label">Deprecated</p>
        <p>For existing framework projects, the build setting appears only after you add a documentation catalog. You
          can’t use the build setting with a Swift package.</p>
      </aside>

      <aside aria-label="experiment" class="experiment">
        <p class="label">Experiment</p>
        <p>For existing framework projects, the build setting appears only after you add a documentation catalog. You
          can’t use the build setting with a Swift package.</p>
      </aside>

      <aside aria-label="tip" class="tip">
        <p class="label">Tip</p>
        <p>For existing framework projects, the build setting appears only after you add a documentation catalog. You
          can’t use the build setting with a Swift package.</p>
      </aside>
    </div>

    <section id="topics" class="contenttable alt-light" data-v-0be2ddac="" data-v-134927b9="" data-v-8dda19e8="">
      <div class="container" data-v-0be2ddac=""><h2 class="title" data-v-0be2ddac="">Topics</h2>
        <div class="row contenttable-section" data-v-1e52db7a="" data-v-5a17d4c0="" data-v-8dda19e8="">
          <div class="col section-title large-3 medium-3 small-12" data-v-1e52db7a="" data-v-3551f3f5=""
               data-v-5a17d4c0=""><h3 class="title" data-v-1e52db7a="" data-v-3551f3f5="">Essentials</h3></div>
          <div class="col section-content large-9 medium-9 small-12" data-v-1e52db7a="" data-v-3551f3f5=""
               data-v-5a17d4c0=""><!----><!----><!---->
            <div class="link-block topic" data-v-13b1fa63="" data-v-3551f3f5="" data-v-8dda19e8=""><a
                class="link has-adjacent-elements"
                data-v-13b1fa63=""
                href="/documentation/xcode/about-continuous-integration-and-delivery-with-xcode-cloud">
              <div class="topic-icon-wrapper" data-v-13b1fa63="" data-v-3e592fb4="">
                <svg aria-hidden="true" class="svg-icon article-icon topic-icon" data-v-3e592fb4="" data-v-7abeccde=""
                     viewBox="0 0 15 15" xmlns="http://www.w3.org/2000/svg">
                  <path
                      d="m11.2623182 15c1.365556 0 2.055373-.7038949 2.055373-2.069451v-6.60957293c0-.76020648-.0985453-1.09103707-.5631159-1.5696856l-4.13890191-4.18113561c-.45049272-.45753168-.81651806-.57015486-1.4992961-.57015486h-3.37869545c-1.35147818 0-2.05537306.70389489-2.05537306 2.07648991v10.85405909c0 1.3725951.69685593 2.069451 2.05537306 2.069451zm-.0422337-.8657907h-7.44016897c-.80947912 0-1.2247771-.4293759-1.2247771-1.2247771v-10.81182544c0-.78132332.41529798-1.2247771 1.23181605-1.2247771h3.28718911v4.23744721c0 .80244016.40825904 1.1825434 1.18958236 1.1825434h4.18817455v6.61661193c0 .7954012-.4223369 1.2247771-1.231816 1.2247771zm1.0558423-8.66494604h-3.92069451c-.32379165 0-.45753168-.12670108-.45753168-.45753168v-3.9629282z"
                      data-v-7abeccde=""></path>
                </svg>
              </div>
              <span data-v-13b1fa63="">About Continuous Integration and Delivery with Xcode Cloud</span><!----></a>
              <div class="abstract" data-v-13b1fa63="">
                <div class="content" data-v-13b1fa63="" data-v-5b9db162="">Learn how continuous integration and delivery
                  with Xcode Cloud helps you create high-quality apps and frameworks.
                </div><!----><!----><!----></div><!----></div>
            <div class="link-block topic" data-v-13b1fa63="" data-v-3551f3f5="" data-v-8dda19e8=""><a
                class="link has-adjacent-elements"
                data-v-13b1fa63=""
                href="/documentation/xcode/requirements-for-using-xcode-cloud">
              <div class="topic-icon-wrapper" data-v-13b1fa63="" data-v-3e592fb4="">
                <svg aria-hidden="true" class="svg-icon article-icon topic-icon" data-v-3e592fb4="" data-v-7abeccde=""
                     viewBox="0 0 15 15" xmlns="http://www.w3.org/2000/svg">
                  <path
                      d="m11.2623182 15c1.365556 0 2.055373-.7038949 2.055373-2.069451v-6.60957293c0-.76020648-.0985453-1.09103707-.5631159-1.5696856l-4.13890191-4.18113561c-.45049272-.45753168-.81651806-.57015486-1.4992961-.57015486h-3.37869545c-1.35147818 0-2.05537306.70389489-2.05537306 2.07648991v10.85405909c0 1.3725951.69685593 2.069451 2.05537306 2.069451zm-.0422337-.8657907h-7.44016897c-.80947912 0-1.2247771-.4293759-1.2247771-1.2247771v-10.81182544c0-.78132332.41529798-1.2247771 1.23181605-1.2247771h3.28718911v4.23744721c0 .80244016.40825904 1.1825434 1.18958236 1.1825434h4.18817455v6.61661193c0 .7954012-.4223369 1.2247771-1.231816 1.2247771zm1.0558423-8.66494604h-3.92069451c-.32379165 0-.45753168-.12670108-.45753168-.45753168v-3.9629282z"
                      data-v-7abeccde=""></path>
                </svg>
              </div>
              <span data-v-13b1fa63="">Requirements for Using Xcode Cloud</span><!----></a>
              <div class="abstract" data-v-13b1fa63="">
                <div class="content" data-v-13b1fa63="" data-v-5b9db162="">Review account, project, and source control
                  requirements before configuring your project or workspace to use Xcode Cloud.
                </div><!----><!----><!----></div><!----></div>
          </div>
        </div>
        <div class="row contenttable-section" data-v-1e52db7a="" data-v-5a17d4c0="" data-v-8dda19e8="">
          <div class="col section-title large-3 medium-3 small-12" data-v-1e52db7a="" data-v-3551f3f5=""
               data-v-5a17d4c0=""><h3 class="title" data-v-1e52db7a="" data-v-3551f3f5="">Setup and Maintenance</h3>
          </div>
          <div class="col section-content large-9 medium-9 small-12" data-v-1e52db7a="" data-v-3551f3f5=""
               data-v-5a17d4c0=""><!----><!----><!---->
            <div class="link-block topic" data-v-13b1fa63="" data-v-3551f3f5="" data-v-8dda19e8=""><a
                class="link has-adjacent-elements"
                data-v-13b1fa63=""
                href="/documentation/xcode/configuring-your-first-xcode-cloud-workflow">
              <div class="topic-icon-wrapper" data-v-13b1fa63="" data-v-3e592fb4="">
                <svg aria-hidden="true" class="svg-icon article-icon topic-icon" data-v-3e592fb4="" data-v-7abeccde=""
                     viewBox="0 0 15 15" xmlns="http://www.w3.org/2000/svg">
                  <path
                      d="m11.2623182 15c1.365556 0 2.055373-.7038949 2.055373-2.069451v-6.60957293c0-.76020648-.0985453-1.09103707-.5631159-1.5696856l-4.13890191-4.18113561c-.45049272-.45753168-.81651806-.57015486-1.4992961-.57015486h-3.37869545c-1.35147818 0-2.05537306.70389489-2.05537306 2.07648991v10.85405909c0 1.3725951.69685593 2.069451 2.05537306 2.069451zm-.0422337-.8657907h-7.44016897c-.80947912 0-1.2247771-.4293759-1.2247771-1.2247771v-10.81182544c0-.78132332.41529798-1.2247771 1.23181605-1.2247771h3.28718911v4.23744721c0 .80244016.40825904 1.1825434 1.18958236 1.1825434h4.18817455v6.61661193c0 .7954012-.4223369 1.2247771-1.231816 1.2247771zm1.0558423-8.66494604h-3.92069451c-.32379165 0-.45753168-.12670108-.45753168-.45753168v-3.9629282z"
                      data-v-7abeccde=""></path>
                </svg>
              </div>
              <span data-v-13b1fa63="">Configuring Your First Xcode Cloud Workflow</span><!----></a>
              <div class="abstract" data-v-13b1fa63="">
                <div class="content" data-v-13b1fa63="" data-v-5b9db162="">Configure your project or workspace to use
                  Xcode Cloud and adopt continuous integration and delivery.
                </div><!----><!----><!----></div><!----></div>
            <div class="link-block topic" data-v-13b1fa63="" data-v-3551f3f5="" data-v-8dda19e8=""><a
                class="link has-adjacent-elements"
                data-v-13b1fa63=""
                href="/documentation/xcode/making-dependencies-available-to-xcode-cloud">
              <div class="topic-icon-wrapper" data-v-13b1fa63="" data-v-3e592fb4="">
                <svg aria-hidden="true" class="svg-icon article-icon topic-icon" data-v-3e592fb4="" data-v-7abeccde=""
                     viewBox="0 0 15 15" xmlns="http://www.w3.org/2000/svg">
                  <path
                      d="m11.2623182 15c1.365556 0 2.055373-.7038949 2.055373-2.069451v-6.60957293c0-.76020648-.0985453-1.09103707-.5631159-1.5696856l-4.13890191-4.18113561c-.45049272-.45753168-.81651806-.57015486-1.4992961-.57015486h-3.37869545c-1.35147818 0-2.05537306.70389489-2.05537306 2.07648991v10.85405909c0 1.3725951.69685593 2.069451 2.05537306 2.069451zm-.0422337-.8657907h-7.44016897c-.80947912 0-1.2247771-.4293759-1.2247771-1.2247771v-10.81182544c0-.78132332.41529798-1.2247771 1.23181605-1.2247771h3.28718911v4.23744721c0 .80244016.40825904 1.1825434 1.18958236 1.1825434h4.18817455v6.61661193c0 .7954012-.4223369 1.2247771-1.231816 1.2247771zm1.0558423-8.66494604h-3.92069451c-.32379165 0-.45753168-.12670108-.45753168-.45753168v-3.9629282z"
                      data-v-7abeccde=""></path>
                </svg>
              </div>
              <span data-v-13b1fa63="">Making Dependencies Available to Xcode Cloud</span><!----></a>
              <div class="abstract" data-v-13b1fa63="">
                <div class="content" data-v-13b1fa63="" data-v-5b9db162="">Review dependencies and make them available
                  to Xcode Cloud before you configure your project to use Xcode Cloud.
                </div><!----><!----><!----></div><!----></div>
            <div class="link-block topic" data-v-13b1fa63="" data-v-3551f3f5="" data-v-8dda19e8=""><a
                class="link has-adjacent-elements"
                data-v-13b1fa63=""
                href="/documentation/xcode/configuring-xcode-cloud-for-your-team">
              <div class="topic-icon-wrapper" data-v-13b1fa63="" data-v-3e592fb4="">
                <svg aria-hidden="true" class="svg-icon article-icon topic-icon" data-v-3e592fb4="" data-v-7abeccde=""
                     viewBox="0 0 15 15" xmlns="http://www.w3.org/2000/svg">
                  <path
                      d="m11.2623182 15c1.365556 0 2.055373-.7038949 2.055373-2.069451v-6.60957293c0-.76020648-.0985453-1.09103707-.5631159-1.5696856l-4.13890191-4.18113561c-.45049272-.45753168-.81651806-.57015486-1.4992961-.57015486h-3.37869545c-1.35147818 0-2.05537306.70389489-2.05537306 2.07648991v10.85405909c0 1.3725951.69685593 2.069451 2.05537306 2.069451zm-.0422337-.8657907h-7.44016897c-.80947912 0-1.2247771-.4293759-1.2247771-1.2247771v-10.81182544c0-.78132332.41529798-1.2247771 1.23181605-1.2247771h3.28718911v4.23744721c0 .80244016.40825904 1.1825434 1.18958236 1.1825434h4.18817455v6.61661193c0 .7954012-.4223369 1.2247771-1.231816 1.2247771zm1.0558423-8.66494604h-3.92069451c-.32379165 0-.45753168-.12670108-.45753168-.45753168v-3.9629282z"
                      data-v-7abeccde=""></path>
                </svg>
              </div>
              <span data-v-13b1fa63="">Configuring Xcode Cloud for Your Team</span><!----></a>
              <div class="abstract" data-v-13b1fa63="">
                <div class="content" data-v-13b1fa63="" data-v-5b9db162="">Start using continuous integration and
                  delivery with Xcode Cloud as a team.
                </div><!----><!----><!----></div><!----></div>
            <div class="link-block topic" data-v-13b1fa63="" data-v-3551f3f5="" data-v-8dda19e8=""><a
                class="link has-adjacent-elements"
                data-v-13b1fa63=""
                href="/documentation/xcode/source-code-management-setup">
              <div class="topic-icon-wrapper" data-v-13b1fa63="" data-v-3e592fb4="">
                <svg aria-hidden="true" class="svg-icon api-reference-icon topic-icon" data-v-3e592fb4=""
                     data-v-7abeccde=""
                     viewBox="0 0 15 15" xmlns="http://www.w3.org/2000/svg"><title data-v-7abeccde="">API
                  Reference</title>
                  <circle cx="3.19" cy="4.49" data-v-7abeccde="" r=".5"></circle>
                  <circle cx="3.19" cy="7.5" data-v-7abeccde="" r=".5"></circle>
                  <circle cx="3.19" cy="10.48" data-v-7abeccde="" r=".5"></circle>
                  <g data-v-7abeccde="" fill-rule="evenodd">
                    <path d="m4.72 7h6.78a.5.5 0 0 1 .5.5.5.5 0 0 1 -.5.5h-6.78a.51.51 0 0 1 -.5-.5.5.5 0 0 1 .5-.5z"
                          data-v-7abeccde=""></path>
                    <path d="m4.72 10h6.78a.5.5 0 0 1 .5.5.5.5 0 0 1 -.5.5h-6.78a.5.5 0 0 1 -.5-.5.51.51 0 0 1 .5-.5z"
                          data-v-7abeccde=""></path>
                    <path d="m4.72 4h6.78a.5.5 0 0 1 .5.5.5.5 0 0 1 -.5.5h-6.78a.5.5 0 0 1 -.5-.5.51.51 0 0 1 .5-.5z"
                          data-v-7abeccde=""></path>
                  </g>
                  <path
                      d="m13 0h-11a2 2 0 0 0 -2 2v11a2 2 0 0 0 2 2h11a2 2 0 0 0 2-2v-11a2 2 0 0 0 -2-2zm1 13a1 1 0 0 1 -1 1h-11a1 1 0 0 1 -1-1v-11a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1z"
                      data-v-7abeccde=""></path>
                </svg>
              </div>
              <span data-v-13b1fa63="">Source Code Management Setup</span><!----></a>
              <div class="abstract" data-v-13b1fa63="">
                <div class="content" data-v-13b1fa63="" data-v-5b9db162="">Allow Xcode Cloud to access your Git
                  repository.
                </div><!----><!----><!----></div><!----></div>
            <div class="link-block topic" data-v-13b1fa63="" data-v-3551f3f5="" data-v-8dda19e8=""><a
                class="link has-adjacent-elements"
                data-v-13b1fa63=""
                href="/documentation/xcode/removing-your-project-from-xcode-cloud">
              <div class="topic-icon-wrapper" data-v-13b1fa63="" data-v-3e592fb4="">
                <svg aria-hidden="true" class="svg-icon article-icon topic-icon" data-v-3e592fb4="" data-v-7abeccde=""
                     viewBox="0 0 15 15" xmlns="http://www.w3.org/2000/svg">
                  <path
                      d="m11.2623182 15c1.365556 0 2.055373-.7038949 2.055373-2.069451v-6.60957293c0-.76020648-.0985453-1.09103707-.5631159-1.5696856l-4.13890191-4.18113561c-.45049272-.45753168-.81651806-.57015486-1.4992961-.57015486h-3.37869545c-1.35147818 0-2.05537306.70389489-2.05537306 2.07648991v10.85405909c0 1.3725951.69685593 2.069451 2.05537306 2.069451zm-.0422337-.8657907h-7.44016897c-.80947912 0-1.2247771-.4293759-1.2247771-1.2247771v-10.81182544c0-.78132332.41529798-1.2247771 1.23181605-1.2247771h3.28718911v4.23744721c0 .80244016.40825904 1.1825434 1.18958236 1.1825434h4.18817455v6.61661193c0 .7954012-.4223369 1.2247771-1.231816 1.2247771zm1.0558423-8.66494604h-3.92069451c-.32379165 0-.45753168-.12670108-.45753168-.45753168v-3.9629282z"
                      data-v-7abeccde=""></path>
                </svg>
              </div>
              <span data-v-13b1fa63="">Removing Your Project from Xcode Cloud</span><!----></a>
              <div class="abstract" data-v-13b1fa63="">
                <div class="content" data-v-13b1fa63="" data-v-5b9db162="">Remove your project from Xcode Cloud to
                  delete app and workflow data, disconnect your Git repository, and remove the Slack integration.
                </div><!----><!----><!----></div><!----></div>
          </div>
        </div>
        <div class="row contenttable-section" data-v-1e52db7a="" data-v-5a17d4c0="" data-v-8dda19e8="">
          <div class="col section-title large-3 medium-3 small-12" data-v-1e52db7a="" data-v-3551f3f5=""
               data-v-5a17d4c0=""><h3 class="title" data-v-1e52db7a="" data-v-3551f3f5="">Workflows</h3></div>
          <div class="col section-content large-9 medium-9 small-12" data-v-1e52db7a="" data-v-3551f3f5=""
               data-v-5a17d4c0=""><!----><!----><!---->
            <div class="link-block topic" data-v-13b1fa63="" data-v-3551f3f5="" data-v-8dda19e8=""><a
                class="link has-adjacent-elements"
                data-v-13b1fa63=""
                href="/documentation/xcode/developing-a-workflow-strategy-for-xcode-cloud">
              <div class="topic-icon-wrapper" data-v-13b1fa63="" data-v-3e592fb4="">
                <svg aria-hidden="true" class="svg-icon article-icon topic-icon" data-v-3e592fb4="" data-v-7abeccde=""
                     viewBox="0 0 15 15" xmlns="http://www.w3.org/2000/svg">
                  <path
                      d="m11.2623182 15c1.365556 0 2.055373-.7038949 2.055373-2.069451v-6.60957293c0-.76020648-.0985453-1.09103707-.5631159-1.5696856l-4.13890191-4.18113561c-.45049272-.45753168-.81651806-.57015486-1.4992961-.57015486h-3.37869545c-1.35147818 0-2.05537306.70389489-2.05537306 2.07648991v10.85405909c0 1.3725951.69685593 2.069451 2.05537306 2.069451zm-.0422337-.8657907h-7.44016897c-.80947912 0-1.2247771-.4293759-1.2247771-1.2247771v-10.81182544c0-.78132332.41529798-1.2247771 1.23181605-1.2247771h3.28718911v4.23744721c0 .80244016.40825904 1.1825434 1.18958236 1.1825434h4.18817455v6.61661193c0 .7954012-.4223369 1.2247771-1.231816 1.2247771zm1.0558423-8.66494604h-3.92069451c-.32379165 0-.45753168-.12670108-.45753168-.45753168v-3.9629282z"
                      data-v-7abeccde=""></path>
                </svg>
              </div>
              <span data-v-13b1fa63="">Developing a Workflow Strategy for Xcode Cloud</span><!----></a>
              <div class="abstract" data-v-13b1fa63="">
                <div class="content" data-v-13b1fa63="" data-v-5b9db162="">Review how you can best create custom Xcode
                  Cloud workflows to refine your continuous integration and delivery practice.
                </div><!----><!----><!----></div><!----></div>
            <div class="link-block topic" data-v-13b1fa63="" data-v-3551f3f5="" data-v-8dda19e8=""><a
                class="link has-adjacent-elements"
                data-v-13b1fa63=""
                href="/documentation/xcode/xcode-cloud-workflow-reference">
              <div class="topic-icon-wrapper" data-v-13b1fa63="" data-v-3e592fb4="">
                <svg aria-hidden="true" class="svg-icon api-reference-icon topic-icon" data-v-3e592fb4=""
                     data-v-7abeccde=""
                     viewBox="0 0 15 15" xmlns="http://www.w3.org/2000/svg"><title data-v-7abeccde="">API
                  Reference</title>
                  <circle cx="3.19" cy="4.49" data-v-7abeccde="" r=".5"></circle>
                  <circle cx="3.19" cy="7.5" data-v-7abeccde="" r=".5"></circle>
                  <circle cx="3.19" cy="10.48" data-v-7abeccde="" r=".5"></circle>
                  <g data-v-7abeccde="" fill-rule="evenodd">
                    <path d="m4.72 7h6.78a.5.5 0 0 1 .5.5.5.5 0 0 1 -.5.5h-6.78a.51.51 0 0 1 -.5-.5.5.5 0 0 1 .5-.5z"
                          data-v-7abeccde=""></path>
                    <path d="m4.72 10h6.78a.5.5 0 0 1 .5.5.5.5 0 0 1 -.5.5h-6.78a.5.5 0 0 1 -.5-.5.51.51 0 0 1 .5-.5z"
                          data-v-7abeccde=""></path>
                    <path d="m4.72 4h6.78a.5.5 0 0 1 .5.5.5.5 0 0 1 -.5.5h-6.78a.5.5 0 0 1 -.5-.5.51.51 0 0 1 .5-.5z"
                          data-v-7abeccde=""></path>
                  </g>
                  <path
                      d="m13 0h-11a2 2 0 0 0 -2 2v11a2 2 0 0 0 2 2h11a2 2 0 0 0 2-2v-11a2 2 0 0 0 -2-2zm1 13a1 1 0 0 1 -1 1h-11a1 1 0 0 1 -1-1v-11a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1z"
                      data-v-7abeccde=""></path>
                </svg>
              </div>
              <span data-v-13b1fa63="">Xcode Cloud Workflow Reference</span><!----></a>
              <div class="abstract" data-v-13b1fa63="">
                <div class="content" data-v-13b1fa63="" data-v-5b9db162="">Configure metadata, start conditions,
                  actions, post-actions, and more to create custom Xcode Cloud workflows.
                </div><!----><!----><!----></div><!----></div>
          </div>
        </div>
        <div class="row contenttable-section" data-v-1e52db7a="" data-v-5a17d4c0="" data-v-8dda19e8="">
          <div class="col section-title large-3 medium-3 small-12" data-v-1e52db7a="" data-v-3551f3f5=""
               data-v-5a17d4c0=""><h3 class="title" data-v-1e52db7a="" data-v-3551f3f5="">Custom Build Scripts</h3>
          </div>
          <div class="col section-content large-9 medium-9 small-12" data-v-1e52db7a="" data-v-3551f3f5=""
               data-v-5a17d4c0=""><!----><!----><!---->
            <div class="link-block topic" data-v-13b1fa63="" data-v-3551f3f5="" data-v-8dda19e8=""><a
                class="link has-adjacent-elements"
                data-v-13b1fa63=""
                href="/documentation/xcode/writing-custom-build-scripts">
              <div class="topic-icon-wrapper" data-v-13b1fa63="" data-v-3e592fb4="">
                <svg aria-hidden="true" class="svg-icon article-icon topic-icon" data-v-3e592fb4="" data-v-7abeccde=""
                     viewBox="0 0 15 15" xmlns="http://www.w3.org/2000/svg">
                  <path
                      d="m11.2623182 15c1.365556 0 2.055373-.7038949 2.055373-2.069451v-6.60957293c0-.76020648-.0985453-1.09103707-.5631159-1.5696856l-4.13890191-4.18113561c-.45049272-.45753168-.81651806-.57015486-1.4992961-.57015486h-3.37869545c-1.35147818 0-2.05537306.70389489-2.05537306 2.07648991v10.85405909c0 1.3725951.69685593 2.069451 2.05537306 2.069451zm-.0422337-.8657907h-7.44016897c-.80947912 0-1.2247771-.4293759-1.2247771-1.2247771v-10.81182544c0-.78132332.41529798-1.2247771 1.23181605-1.2247771h3.28718911v4.23744721c0 .80244016.40825904 1.1825434 1.18958236 1.1825434h4.18817455v6.61661193c0 .7954012-.4223369 1.2247771-1.231816 1.2247771zm1.0558423-8.66494604h-3.92069451c-.32379165 0-.45753168-.12670108-.45753168-.45753168v-3.9629282z"
                      data-v-7abeccde=""></path>
                </svg>
              </div>
              <span data-v-13b1fa63="">Writing Custom Build Scripts</span><!----></a>
              <div class="abstract" data-v-13b1fa63="">
                <div class="content" data-v-13b1fa63="" data-v-5b9db162="">Extend your Xcode Cloud workflows with custom
                  build scripts that perform custom tasks or install additional tools.
                </div><!----><!----><!----></div><!----></div>
            <div class="link-block topic" data-v-13b1fa63="" data-v-3551f3f5="" data-v-8dda19e8=""><a
                class="link has-adjacent-elements"
                data-v-13b1fa63=""
                href="/documentation/xcode/environment-variable-reference">
              <div class="topic-icon-wrapper" data-v-13b1fa63="" data-v-3e592fb4="">
                <svg aria-hidden="true" class="svg-icon article-icon topic-icon" data-v-3e592fb4="" data-v-7abeccde=""
                     viewBox="0 0 15 15" xmlns="http://www.w3.org/2000/svg">
                  <path
                      d="m11.2623182 15c1.365556 0 2.055373-.7038949 2.055373-2.069451v-6.60957293c0-.76020648-.0985453-1.09103707-.5631159-1.5696856l-4.13890191-4.18113561c-.45049272-.45753168-.81651806-.57015486-1.4992961-.57015486h-3.37869545c-1.35147818 0-2.05537306.70389489-2.05537306 2.07648991v10.85405909c0 1.3725951.69685593 2.069451 2.05537306 2.069451zm-.0422337-.8657907h-7.44016897c-.80947912 0-1.2247771-.4293759-1.2247771-1.2247771v-10.81182544c0-.78132332.41529798-1.2247771 1.23181605-1.2247771h3.28718911v4.23744721c0 .80244016.40825904 1.1825434 1.18958236 1.1825434h4.18817455v6.61661193c0 .7954012-.4223369 1.2247771-1.231816 1.2247771zm1.0558423-8.66494604h-3.92069451c-.32379165 0-.45753168-.12670108-.45753168-.45753168v-3.9629282z"
                      data-v-7abeccde=""></path>
                </svg>
              </div>
              <span data-v-13b1fa63="">Environment Variable Reference</span><!----></a>
              <div class="abstract" data-v-13b1fa63="">
                <div class="content" data-v-13b1fa63="" data-v-5b9db162="">Review predefined environment variables you
                  use in custom build scripts.
                </div><!----><!----><!----></div><!----></div>
          </div>
        </div>
        <div class="row contenttable-section" data-v-1e52db7a="" data-v-5a17d4c0="" data-v-8dda19e8="">
          <div class="col section-title large-3 medium-3 small-12" data-v-1e52db7a="" data-v-3551f3f5=""
               data-v-5a17d4c0=""><h3 class="title" data-v-1e52db7a="" data-v-3551f3f5="">Troubleshooting</h3></div>
          <div class="col section-content large-9 medium-9 small-12" data-v-1e52db7a="" data-v-3551f3f5=""
               data-v-5a17d4c0=""><!----><!----><!---->
            <div class="link-block topic" data-v-13b1fa63="" data-v-3551f3f5="" data-v-8dda19e8=""><a
                class="link has-adjacent-elements"
                data-v-13b1fa63=""
                href="/documentation/xcode/resolving-common-configuration-and-build-issues">
              <div class="topic-icon-wrapper" data-v-13b1fa63="" data-v-3e592fb4="">
                <svg aria-hidden="true" class="svg-icon article-icon topic-icon" data-v-3e592fb4="" data-v-7abeccde=""
                     viewBox="0 0 15 15" xmlns="http://www.w3.org/2000/svg">
                  <path
                      d="m11.2623182 15c1.365556 0 2.055373-.7038949 2.055373-2.069451v-6.60957293c0-.76020648-.0985453-1.09103707-.5631159-1.5696856l-4.13890191-4.18113561c-.45049272-.45753168-.81651806-.57015486-1.4992961-.57015486h-3.37869545c-1.35147818 0-2.05537306.70389489-2.05537306 2.07648991v10.85405909c0 1.3725951.69685593 2.069451 2.05537306 2.069451zm-.0422337-.8657907h-7.44016897c-.80947912 0-1.2247771-.4293759-1.2247771-1.2247771v-10.81182544c0-.78132332.41529798-1.2247771 1.23181605-1.2247771h3.28718911v4.23744721c0 .80244016.40825904 1.1825434 1.18958236 1.1825434h4.18817455v6.61661193c0 .7954012-.4223369 1.2247771-1.231816 1.2247771zm1.0558423-8.66494604h-3.92069451c-.32379165 0-.45753168-.12670108-.45753168-.45753168v-3.9629282z"
                      data-v-7abeccde=""></path>
                </svg>
              </div>
              <span data-v-13b1fa63="">Resolving Common Configuration and Build Issues</span><!----></a>
              <div class="abstract" data-v-13b1fa63="">
                <div class="content" data-v-13b1fa63="" data-v-5b9db162="">Review common configuration and build issues
                  and learn how you can resolve them.
                </div><!----><!----><!----></div><!----></div>
          </div>
        </div>
        <div class="row contenttable-section" data-v-1e52db7a="" data-v-5a17d4c0="" data-v-8dda19e8="">
          <div class="col section-title large-3 medium-3 small-12" data-v-1e52db7a="" data-v-3551f3f5=""
               data-v-5a17d4c0=""><h3 class="title" data-v-1e52db7a="" data-v-3551f3f5="">Notifications</h3></div>
          <div class="col section-content large-9 medium-9 small-12" data-v-1e52db7a="" data-v-3551f3f5=""
               data-v-5a17d4c0=""><!----><!----><!---->
            <div class="link-block topic" data-v-13b1fa63="" data-v-3551f3f5="" data-v-8dda19e8=""><a
                class="link has-adjacent-elements"
                data-v-13b1fa63=""
                href="/documentation/xcode/configuring-webhooks-in-xcode-cloud">
              <div class="topic-icon-wrapper" data-v-13b1fa63="" data-v-3e592fb4="">
                <svg aria-hidden="true" class="svg-icon article-icon topic-icon" data-v-3e592fb4="" data-v-7abeccde=""
                     viewBox="0 0 15 15" xmlns="http://www.w3.org/2000/svg">
                  <path
                      d="m11.2623182 15c1.365556 0 2.055373-.7038949 2.055373-2.069451v-6.60957293c0-.76020648-.0985453-1.09103707-.5631159-1.5696856l-4.13890191-4.18113561c-.45049272-.45753168-.81651806-.57015486-1.4992961-.57015486h-3.37869545c-1.35147818 0-2.05537306.70389489-2.05537306 2.07648991v10.85405909c0 1.3725951.69685593 2.069451 2.05537306 2.069451zm-.0422337-.8657907h-7.44016897c-.80947912 0-1.2247771-.4293759-1.2247771-1.2247771v-10.81182544c0-.78132332.41529798-1.2247771 1.23181605-1.2247771h3.28718911v4.23744721c0 .80244016.40825904 1.1825434 1.18958236 1.1825434h4.18817455v6.61661193c0 .7954012-.4223369 1.2247771-1.231816 1.2247771zm1.0558423-8.66494604h-3.92069451c-.32379165 0-.45753168-.12670108-.45753168-.45753168v-3.9629282z"
                      data-v-7abeccde=""></path>
                </svg>
              </div>
              <span data-v-13b1fa63="">Configuring Webhooks in Xcode Cloud</span><!----></a>
              <div class="abstract" data-v-13b1fa63="">
                <div class="content" data-v-13b1fa63="" data-v-5b9db162="">Configure webhooks that connect Xcode Cloud
                  to other services and tools.
                </div><!----><!----><!----></div><!----></div>
            <div class="link-block topic" data-v-13b1fa63="" data-v-3551f3f5="" data-v-8dda19e8=""><a
                class="link has-adjacent-elements"
                data-v-13b1fa63=""
                href="/documentation/xcode/connecting-xcode-cloud-to-slack">
              <div class="topic-icon-wrapper" data-v-13b1fa63="" data-v-3e592fb4="">
                <svg aria-hidden="true" class="svg-icon article-icon topic-icon" data-v-3e592fb4="" data-v-7abeccde=""
                     viewBox="0 0 15 15" xmlns="http://www.w3.org/2000/svg">
                  <path
                      d="m11.2623182 15c1.365556 0 2.055373-.7038949 2.055373-2.069451v-6.60957293c0-.76020648-.0985453-1.09103707-.5631159-1.5696856l-4.13890191-4.18113561c-.45049272-.45753168-.81651806-.57015486-1.4992961-.57015486h-3.37869545c-1.35147818 0-2.05537306.70389489-2.05537306 2.07648991v10.85405909c0 1.3725951.69685593 2.069451 2.05537306 2.069451zm-.0422337-.8657907h-7.44016897c-.80947912 0-1.2247771-.4293759-1.2247771-1.2247771v-10.81182544c0-.78132332.41529798-1.2247771 1.23181605-1.2247771h3.28718911v4.23744721c0 .80244016.40825904 1.1825434 1.18958236 1.1825434h4.18817455v6.61661193c0 .7954012-.4223369 1.2247771-1.231816 1.2247771zm1.0558423-8.66494604h-3.92069451c-.32379165 0-.45753168-.12670108-.45753168-.45753168v-3.9629282z"
                      data-v-7abeccde=""></path>
                </svg>
              </div>
              <span data-v-13b1fa63="">Connecting Xcode Cloud to Slack</span><!----></a>
              <div class="abstract" data-v-13b1fa63="">
                <div class="content" data-v-13b1fa63="" data-v-5b9db162="">Connect Xcode Cloud to Slack to keep your
                  team informed about the latest Xcode Cloud builds.
                </div><!----><!----><!----></div><!----></div>
          </div>
        </div>
        <div class="row contenttable-section" data-v-1e52db7a="" data-v-5a17d4c0="" data-v-8dda19e8="">
          <div class="col section-title large-3 medium-3 small-12" data-v-1e52db7a="" data-v-3551f3f5=""
               data-v-5a17d4c0=""><h3 class="title" data-v-1e52db7a="" data-v-3551f3f5="">REST API</h3></div>
          <div class="col section-content large-9 medium-9 small-12" data-v-1e52db7a="" data-v-3551f3f5=""
               data-v-5a17d4c0=""><!----><!----><!---->
            <div class="link-block topic" data-v-13b1fa63="" data-v-3551f3f5="" data-v-8dda19e8=""><a
                class="link has-adjacent-elements"
                data-v-13b1fa63=""
                href="/documentation/appstoreconnectapi/xcode_cloud_workflows_and_builds">
              <div class="topic-icon-wrapper" data-v-13b1fa63="" data-v-3e592fb4="">
                <svg aria-hidden="true" class="svg-icon article-icon topic-icon" data-v-3e592fb4="" data-v-7abeccde=""
                     viewBox="0 0 15 15" xmlns="http://www.w3.org/2000/svg">
                  <path
                      d="m11.2623182 15c1.365556 0 2.055373-.7038949 2.055373-2.069451v-6.60957293c0-.76020648-.0985453-1.09103707-.5631159-1.5696856l-4.13890191-4.18113561c-.45049272-.45753168-.81651806-.57015486-1.4992961-.57015486h-3.37869545c-1.35147818 0-2.05537306.70389489-2.05537306 2.07648991v10.85405909c0 1.3725951.69685593 2.069451 2.05537306 2.069451zm-.0422337-.8657907h-7.44016897c-.80947912 0-1.2247771-.4293759-1.2247771-1.2247771v-10.81182544c0-.78132332.41529798-1.2247771 1.23181605-1.2247771h3.28718911v4.23744721c0 .80244016.40825904 1.1825434 1.18958236 1.1825434h4.18817455v6.61661193c0 .7954012-.4223369 1.2247771-1.231816 1.2247771zm1.0558423-8.66494604h-3.92069451c-.32379165 0-.45753168-.12670108-.45753168-.45753168v-3.9629282z"
                      data-v-7abeccde=""></path>
                </svg>
              </div>
              <span data-v-13b1fa63="">Xcode Cloud Workflows and Builds</span><!----></a>
              <div class="abstract" data-v-13b1fa63="">
                <div class="content" data-v-13b1fa63="" data-v-5b9db162="">Automate reading Xcode Cloud data, managing
                  workflows, and starting builds.
                </div><!----><!----><!----></div><!----></div>
          </div>
        </div>
      </div>
    </section>

  </div>
</template>
