<script setup>

</script>

<template>
  <nav class="bg-gray-700 py-3 text-base font-normal mb-8">
    <ul aria-label="Breadcrumbs" class="flex list-none lg:max-w-5xl mx-auto px-10">
      <li class="inline-flex">


        <a class="text-gray-50" href="/">docgo Documentation</a>
      </li>

      <li class="inline-flex">
        <div class="w-3 m-2 flex fill-current text-gray-500">
          <svg aria-hidden="true" data-v-134594fd="" data-v-838665fe="" viewBox="0 0 14 14"
               xmlns="http://www.w3.org/2000/svg">
            <path
                d="m4.81347656 13.1269531c.22558594 0 .41015625-.0820312.56738282-.2324219l5.31835942-5.19531245c.1845703-.19140625.2802734-.38964844.2802734-.63574219 0-.23925781-.0888672-.45117187-.2802734-.62890625l-5.31835942-5.20214843c-.15722657-.15039063-.34179688-.23242188-.56738282-.23242188-.45800781 0-.81347656.35546875-.81347656.80664062 0 .21875.09570312.43066407.24609375.58789063l4.79199219 4.67578125-4.79199219 4.6621094c-.15722656.1640625-.24609375.3623047-.24609375.5878906 0 .4511719.35546875.8066406.81347656.8066406z"
                data-v-7abeccde=""></path>
          </svg>
        </div>


        <a class="text-gray-50" href="/doc-1">Doc 1</a>
      </li>
    </ul>
  </nav>

  <main class="lg:max-w-5xl mx-auto px-10">
    <h1>Documenting a Swift Framework or Package</h1>
    <div class="sm:flex mt-8 mb-16">

      <div class="sm:w-9/12">

        <div class="mb-8 text-xl font-light">Create rich and engaging documentation from your in-source comments, and
          add a
          documentation
          catalog to your project to provide additional content.
        </div>

        <div class="pt-8 border-t">

          <h2>Overview</h2>

          <p>DocC, or Documentation Compiler, makes it easy for you to produce documentation for your Swift frameworks
            and
            packages. The compiler builds your documentation by combining comments you write in-source with extension
            files,
            articles, and other resources that live alongside your project in Xcode, allowing you to create rich and
            engaging documentation for developers.</p>

          <p>With DocC, you provide a combination of reference and conceptual content, and connect it together using
            powerful organization and linking capabilities. The compiler integrates directly with Xcode to enhance
            your
            existing workflow, including code completion, Quick Help, and more. And because you write documentation
            directly
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
              <p>Enter “build documentation” in the search box to locate the Build Documentation during ‘Build’
                setting.</p>
            </li>
            <li>
              <p>Choose Yes from the setting’s pop-up button to enable the build setting.</p>
            </li>
          </ol>

          <img
              src="https://docs-assets.developer.apple.com/published/fdd6b5d865669db0f976405dc24b10ac/11800/doc-viewer-parameters@2x.png">


          <div class="code-listing">
            <pre>123123</pre>
          </div>

          <aside aria-label="note" class="note mt-6 p-4 rounded-2xl bg-gray-50 border border-gray-400 shadow">
            <label class="text-gray-700 font-normal">Note</label>
            <p class="mt-2">For existing framework projects, the build setting appears only after you add a
              documentation catalog.
              You
              can’t use the build setting with a Swift package.</p>
          </aside>

          <aside aria-label="important"
                 class="important mt-6 p-4 rounded-2xl bg-yellow-50 border border-yellow-700 shadow">
            <label class="text-yellow-700 font-normal">Important</label>
            <p class="mt-2">For existing framework projects, the build setting appears only after you add a
              documentation catalog.
              You
              can’t use the build setting with a Swift package.</p>
          </aside>

          <aside aria-label="deprecated"
                 class="deprecated mt-6 p-4 rounded-2xl bg-pink-50 border border-yellow-700 shadow">
            <label class="text-yellow-700 font-normal">Deprecated</label>

            <p class="mt-2">For existing framework projects, the build setting appears only after you add a
              documentation catalog.
              You
              can’t use the build setting with a Swift package.</p>
          </aside>

          <aside aria-label="experiment"
                 class="experiment mt-6 p-4 rounded-2xl bg-purple-50 border border-purple-800 shadow">
            <label class="text-purple-800 font-normal">Experiment</label>
            <p class="mt-2">For existing framework projects, the build setting appears only after you add a
              documentation catalog.
              You
              can’t use the build setting with a Swift package.</p>
          </aside>

          <aside aria-label="tip" class="tip mt-6 p-4 rounded-2xl bg-green-50 border border-green-700 shadow">
            <label class="text-green-700 font-normal">Tip</label>
            <p class="mt-2">For existing framework projects, the build setting appears only after you add a
              documentation catalog.
              You
              can’t use the build setting with a Swift package.</p>
          </aside>


        </div>

      </div>

      <div class="sm:w-3/12 font-medium text-base">On This Page</div>

    </div>
  </main>


  <section class="bg-gray-50 pt-8">
    <div class="lg:max-w-5xl mx-auto px-10 mt-5">
      <h2 class="title">Topics</h2>
      <div class="sm:flex sm:border-t mt-8">
        <div class="sm:w-1/4 pb-4 border-b sm:border-none">
          <h3>Essentials</h3>
        </div>
        <div class="sm:w-3/4 ">
          <div class="mt-8">
            <a class="inline-flex"
               href="/documentation/xcode/about-continuous-integration-and-delivery-with-xcode-cloud">
              <div class="w-4 flex mr-4 text-gray-500 fill-current">
                <svg aria-hidden="true" class=""
                     viewBox="0 0 15 15" xmlns="http://www.w3.org/2000/svg">
                  <path
                      d="m11.2623182 15c1.365556 0 2.055373-.7038949 2.055373-2.069451v-6.60957293c0-.76020648-.0985453-1.09103707-.5631159-1.5696856l-4.13890191-4.18113561c-.45049272-.45753168-.81651806-.57015486-1.4992961-.57015486h-3.37869545c-1.35147818 0-2.05537306.70389489-2.05537306 2.07648991v10.85405909c0 1.3725951.69685593 2.069451 2.05537306 2.069451zm-.0422337-.8657907h-7.44016897c-.80947912 0-1.2247771-.4293759-1.2247771-1.2247771v-10.81182544c0-.78132332.41529798-1.2247771 1.23181605-1.2247771h3.28718911v4.23744721c0 .80244016.40825904 1.1825434 1.18958236 1.1825434h4.18817455v6.61661193c0 .7954012-.4223369 1.2247771-1.231816 1.2247771zm1.0558423-8.66494604h-3.92069451c-.32379165 0-.45753168-.12670108-.45753168-.45753168v-3.9629282z"></path>
                </svg>
              </div>
              <span>About Continuous Integration and Delivery with Xcode Cloud</span>
            </a>
            <div class="ml-8">
              <div class="content" data-v-5b9db162="">Learn how continuous integration and
                delivery
                with Xcode Cloud helps you create high-quality apps and frameworks.
              </div>
            </div>
          </div>
          <div class="mt-8">
            <a class="inline-flex"
               href="/documentation/xcode/about-continuous-integration-and-delivery-with-xcode-cloud">
              <div class="w-4 flex mr-4 text-gray-500 fill-current">
                <svg aria-hidden="true" class=""
                     viewBox="0 0 15 15" xmlns="http://www.w3.org/2000/svg">
                  <path
                      d="m11.2623182 15c1.365556 0 2.055373-.7038949 2.055373-2.069451v-6.60957293c0-.76020648-.0985453-1.09103707-.5631159-1.5696856l-4.13890191-4.18113561c-.45049272-.45753168-.81651806-.57015486-1.4992961-.57015486h-3.37869545c-1.35147818 0-2.05537306.70389489-2.05537306 2.07648991v10.85405909c0 1.3725951.69685593 2.069451 2.05537306 2.069451zm-.0422337-.8657907h-7.44016897c-.80947912 0-1.2247771-.4293759-1.2247771-1.2247771v-10.81182544c0-.78132332.41529798-1.2247771 1.23181605-1.2247771h3.28718911v4.23744721c0 .80244016.40825904 1.1825434 1.18958236 1.1825434h4.18817455v6.61661193c0 .7954012-.4223369 1.2247771-1.231816 1.2247771zm1.0558423-8.66494604h-3.92069451c-.32379165 0-.45753168-.12670108-.45753168-.45753168v-3.9629282z"></path>
                </svg>
              </div>
              <span>About Continuous Integration and Delivery with Xcode Cloud</span>
            </a>
            <div class="ml-8">
              <div class="content" data-v-5b9db162="">Learn how continuous integration and
                delivery
                with Xcode Cloud helps you create high-quality apps and frameworks.
              </div>
            </div>
          </div>

        </div>
      </div>
    </div>

  </section>
</template>
