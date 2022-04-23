# Question:
### I see you are choosing to use external APIs. What advantages and disadvantages does this approach have for your project?
<br>

## Advantages
<ol>
  <li>Less Processing Power.</li>
  <li>Not required to write the algorithm yourself.</li>
  <li>Easy to scale into multiple live services off fewer queries.</li>
</ol>

## Disadvantages
<ol>
  <li>Reliant on external algorithms which aren't written by us.</li>
  <li>If those live services all go down then we can only rely on backups.</li>
</ol>

## Recommendations
<ol>
  <li>Allow for multiple services to be queried instead of just one, less of a reliance.</li>
  <li>Create backups of the data to ensure consistency and to cause less spam.</li>
</ol>